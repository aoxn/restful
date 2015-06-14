package apiserver

import (
    "github.com/emicklei/go-restful"
    "net/http"
    "fmt"
    "time"
)


type WebRouter struct {
    h_map map[string]map[string]restful.RouteFunction
}

type NodeServer struct {
}

func NewNodeServer() *NodeServer {
    svr := &NodeServer{}
    svc := restful.WebService{}
    svc.Path("/eggo/").Produces(restful.MIME_JSON).Consumes(restful.MIME_JSON)
    router := newWebRouter(svr)
    func(h *WebRouter) {
        //注册handler
        for method, routes := range router.h_map {
            for route, fct := range routes {
                svc.Route(svc.Method(method).Path(route).To(fct))
            }
        }
    }(router)
    restful.Add(&svc)
    return svr
}

//Start the server
func (s *NodeServer) Start() error {
    //TODO: use container based service loop
    go http.ListenAndServe(":8088", nil)
    return nil
}


func newWebRouter(handler *NodeServer) *WebRouter {
    m := map[string]map[string]restful.RouteFunction{
        "PUT": {
            "nodes/task":handler.CreateTask,
            //"/pkgmaker/pkg/{pkg_id}/files":                          handler.uploadPackageFile,
            //"/images/{name:.*}/get":           s.getImagesGet,
        },
        "GET": {
            "nodes/task":handler.CreateTask,
            //"/pkgmaker/pkg/{pkg_id}/create":                          handler.createInstallPackage,
            //"/images/{name:.*}/get":           s.getImagesGet,
            //"/images/{name:.*}/history":       s.getImagesHistory,
            //"/images/{name:.*}/json":          s.getImagesByName,
            //"/containers/ps":                  s.getContainersJSON,
            //"/containers/json":                s.getContainersJSON,
            //"/containers/{name:.*}/export":    s.getContainersExport,
            //"/containers/{name:.*}/changes":   s.getContainersChanges,

        },
        "POST": {
            "nodes/task":handler.CreateTask,
            //"/pkgmaker/pkg/{pkg_id}/files":                          handler.uploadPackageFile,
            //"/commit":                       s.postCommit,
            //"/build":                        s.postBuild,
            //"/images/create":                s.postImagesCreate,
            //"/images/load":                  s.postImagesLoad,
            //"/images/{name:.*}/push":        s.postImagesPush,
        },
        "DELETE": {
            //"/containers/{name:.*}": s.deleteContainers,
            //"/images/{name:.*}":     s.deleteImages,
        },
        "OPTIONS": {
            //"": s.optionsHandler,
        },
    }
    return &WebRouter{h_map:m}
}

func (h * NodeServer) responseSuccess(request *restful.Request,
response *restful.Response,
courier interface{}) {
    response.StatusCode()
    response.AddHeader("Last-Modified", time.Now().Add(time.Duration(1000)).String())
    response.WriteAsJson(courier)
    return
}

//每个handler中不可以直接使用全局变量，使用前请加锁
func (h * NodeServer) CreateTask(request *restful.Request,
response *restful.Response) {


    fmt.Println("XCreateTask was executed")

    h.responseSuccess(request, response, nil)
    return
}

