function HttpProxy($http) {
    var proxyUrl = "/api/ui/httpproxy";
    return {
        get: function (url, cb) {
            $http.post(proxyUrl, { url: url, method: "GET" }).then(cb);
        },
        post: function(url, data, cb) {
            $http.post(proxyUrl, { url: url, method: "POST", data: data }).then(cb);
        },
        put: function (url, data, cb) {
            $http.post(proxyUrl, { url: url, method: "PUT", data: data }).then(cb);
        },
        delete: function (url, data, cb) {
            $http.post(proxyUrl, { url: url, method: "DELETE", data: data }).then(cb);
        }
    }
}

function ThingManager($http) {
    return {
        create: function(t, cb) {
            $http({
                method: "POST",
                url: "/api/thing",
                data: t
            }).success(cb);
        },
        delete: function (id, cb) {
            $http({
                method: "DELETE",
                url: "/api/thing/" + id
            }).success(cb);
        }
    }
}

// Events

// UI

// Rules