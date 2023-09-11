package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RenanFerreira0023/FiberTemp/config"
	controllerAgent "github.com/RenanFerreira0023/FiberTemp/controllers/agent"
	middlewareController "github.com/RenanFerreira0023/FiberTemp/controllers/middleware"
	controllerReceptor "github.com/RenanFerreira0023/FiberTemp/controllers/receptor"
	repositoryAgent "github.com/RenanFerreira0023/FiberTemp/repositories/agent"
	repositoryReceptor "github.com/RenanFerreira0023/FiberTemp/repositories/receptor"
	"github.com/joho/godotenv"
)

func NewRouter() http.Handler {
	db := config.NewDB()

	agentRepository := repositoryAgent.NewAgentRepository(db)
	agentController := controllerAgent.NewAgentController(agentRepository)

	receptorRepository := repositoryReceptor.NewReceptorRepository(db)
	receptorController := controllerReceptor.NewReceptorController(receptorRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("/Health", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			})).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Repector/Auth/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				receptorController.CheckURLDatas(
					receptorController.CheckUserExist(
						//middlewareController.CreateAuthMiddleware(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Receptor/Create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.InsertReceptor(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Receptor/Login/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.CheckURLDatas(
						receptorController.GetLoginReceptor(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Receptor/List/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.CheckURLDatas(
						receptorController.GetListReceptor(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Receptor/Delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "DELETE":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.DeleteReceptor(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Receptor/Edit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "PUT":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.EditReceptor(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Copy/Find/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.CheckURLDatas(
						receptorController.GetCopy(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	//
	mux.HandleFunc("/Copy/Reply", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.SendReqCopy(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	//
	//
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	/////~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	mux.HandleFunc("/Agent/Auth/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				agentController.CheckURLDatas(
					agentController.CheckUserExist(
						//		middlewareController.CreateAuthMiddleware(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Agent/Create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CreateAgent(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/Create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CreateChannel(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/List/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CheckURLDatas(
						agentController.GetListChannel(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/Permission/List/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CheckURLDatas(
						agentController.GetListPermissionChannel(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/Delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "DELETE":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.DeleteChannel(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/Update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "PUT":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.UpdateChannel(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Agent/Login/Adm", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					//					agentController.CheckURLDatas(
					agentController.GetLoginAgentAdm(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Agent/LoginMT5/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CheckURLDatas(
						agentController.GetLoginAgentMt5(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Copy/Send", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.InsertCopy(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Channel/Permission/Insert", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "OPTIONS":
			handleOptionsRequest(w, r)
			return
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.InsertPermissionChannel(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
		})).ServeHTTP(w, r)
	})

	return mux
}

func handleOptionsRequest(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}
	frontHostURL := os.Getenv("FRONT_HOST_URL")
	w.Header().Set("Access-Control-Allow-Origin", frontHostURL)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
}
