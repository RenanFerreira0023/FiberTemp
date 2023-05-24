package routers

import (
	"net/http"

	"rds_api_2/config"
	controllerAgent "rds_api_2/controllers/agent"
	middlewareController "rds_api_2/controllers/middleware"
	controllerReceptor "rds_api_2/controllers/receptor"
	repositoryAgent "rds_api_2/repositories/agent"
	repositoryReceptor "rds_api_2/repositories/receptor"
)

func NewRouter() http.Handler {
	db := config.NewDB()

	agentRepository := repositoryAgent.NewAgentRepository(db)
	agentController := controllerAgent.NewAgentController(agentRepository)

	receptorRepository := repositoryReceptor.NewReceptorRepository(db)
	receptorController := controllerReceptor.NewReceptorController(receptorRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("/AuthRepector/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				receptorController.CheckURLDatas(
					receptorController.CheckUserExist(
						middlewareController.CreateAuthMiddleware(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/InsertReceptor", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.InsertReceptor(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/LoginReceptor/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.CheckURLDatas(
						receptorController.GetLoginReceptor(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/FindCopy/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.CheckURLDatas(
						receptorController.GetCopy(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

	//
	mux.HandleFunc("/SendReqCopy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					receptorController.SendReqCopy(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

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

	mux.HandleFunc("/AuthAgent/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				agentController.CheckURLDatas(
					agentController.CheckUserExist(
						middlewareController.CreateAuthMiddleware(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/InsertAgent", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.InsertAgent(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/CreateChannel", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CreateChannel(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/LoginAgent/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.CheckURLDatas(
						agentController.GetLoginAgent(
							http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

							}))))).ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/SendCopy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middlewareController.CheckAntiDDoS(
				middlewareController.CheckValidToken(
					agentController.SendCopy(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

						})))).ServeHTTP(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})

	return mux
}
