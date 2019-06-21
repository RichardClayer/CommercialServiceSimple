package routes

import "github.com/BiLuoHui/CommercialServiceSimple/pkg/services"

var PayRoutes = Routes{
    Route{
        "jsapi",
        "post",
        "jsapi",
        services.JSAPIPay,
    },
}