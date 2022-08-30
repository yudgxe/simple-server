package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/stan.go"
	"github.com/yudgxe/simple-server/internal/model"
	"github.com/yudgxe/simple-server/internal/storage"
)

type Server struct {
	config *Config
	store  *storage.Storage
	cache  *model.Order
}

func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() error {

	if err := s.configureStore(); err != nil {
		return err
	}

	if err := s.configureCache(); err != nil {
		return err
	}

	if err := s.configureNats(); err != nil {
		return err
	}
	s.configureRouter()

	log.Println("Server starting")

	return http.ListenAndServe(s.config.BindAddr, nil)
}

func (s *Server) configureRouter() {
	http.HandleFunc("/", s.HandleOrder())
	//http.HandleFunc("/", controller.HandleOrder(s.cache))
}

func (s *Server) configureStore() error {
	st := storage.New(s.config.Storage)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *Server) configureNats() error {
	sc, err := stan.Connect(s.config.Subscriber.ClusterID, s.config.Subscriber.ClientID)
	if err != nil {
		return err
	}

	sc.Subscribe(s.config.Subscriber.Subject, func(m *stan.Msg) {
		order := &model.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Fatal(err)
			return
		}

		if s.cache == nil {
			s.PushData(order)
			s.cache = order
		}

	})

	return nil
}

func (s *Server) configureCache() error {
	cache, err := s.store.Cache().Get()
	if err != nil {
		return err
	}

	if cache != nil {
		s.cache = s.PullData(&cache.OrderID)
	}

	return nil
}

func (s *Server) PushData(o *model.Order) {
	if err := s.store.Order().Create(o); err != nil {
		log.Println(err)
	}

	if err := s.store.Payment().Create(o.Payment); err != nil {
		log.Println(err)
	}

	if err := s.store.Item().Create(o.Items); err != nil {
		log.Println(err)
	}

	if err := s.store.Delivery().Create(o.Delivery, &o.UID); err != nil {
		log.Println(err)
	}

	if err := s.store.Cache().Create(&model.Cache{OrderID: o.UID}); err != nil {
		log.Println(err)
	}
}

func (s *Server) PullData(orderID *string) *model.Order {
	o, err := s.store.Order().FindByID(orderID)
	if err != nil {
		log.Println(err)
	}
	p, err := s.store.Payment().FindByTransaction(orderID)
	if err != nil {
		log.Println(err)
	}

	i, err := s.store.Item().FindByTrackNumber(&o.TrackNumber)
	if err != nil {
		log.Println(err)
	}

	d, err := s.store.Delivery().FindByOrderID(orderID)
	if err != nil {
		log.Println(err)
	}

	o.Payment = p
	o.Items = i
	o.Delivery = d

	return o
}

func (s *Server) HandleOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if s.cache != nil {
			orderID := r.URL.Path[1:]
			if orderID == s.cache.UID {
				res, err := json.Marshal(s.cache)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}
				w.Write(res)
				return
			}
		}
		res, err := json.Marshal(&model.Order{})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(res)
	}
}
