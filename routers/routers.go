package routers

import (
	"net/http"

	"github.com/RenanFerreira0023/FiberTemp/config"
	controllerAgent "github.com/RenanFerreira0023/FiberTemp/controllers/agent"
	middlewareController "github.com/RenanFerreira0023/FiberTemp/controllers/middleware"
	controllerReceptor "github.com/RenanFerreira0023/FiberTemp/controllers/receptor"
	repositoryAgent "github.com/RenanFerreira0023/FiberTemp/repositories/agent"
	repositoryReceptor "github.com/RenanFerreira0023/FiberTemp/repositories/receptor"
)

func NewRouter() http.Handler {
	db := config.NewDB()

	agentRepository := repositoryAgent.NewAgentRepository(db)
	agentController := controllerAgent.NewAgentController(agentRepository)

	receptorRepository := repositoryReceptor.NewReceptorRepository(db)
	receptorController := controllerReceptor.NewReceptorController(receptorRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("/Auth/Repector/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Create/Receptor", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Login/Receptor/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Find/Copy/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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
	mux.HandleFunc("/Send/Reply/Copy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Auth/Agent/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Create/Agent", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.InsertAgent(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Create/Channel", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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

	mux.HandleFunc("/Login/Agent/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CheckURLDatas(
						agentController.GetLoginAgent(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Send/Copy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.SendCopy(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			middlewareController.CheckAntiDDoS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, middlewareController.ConvertStructError(http.StatusText(http.StatusMethodNotAllowed)), http.StatusBadRequest)
			})).ServeHTTP(w, r)

		}
	})

	mux.HandleFunc("/Insert/PermissionChannel", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
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
