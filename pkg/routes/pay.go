package routes

import "github.com/BiLuoHui/CommercialServiceSimple/pkg/services"

var PayRoutes = Routes{
    Route{
        "jsapi",
        "POST",
        "/jsapi",
        services.JSAPIPay,
    },
}