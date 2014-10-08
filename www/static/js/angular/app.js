function HttpProxy($http) {
    var proxyUrl = "/api/ui/proxy";
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
        get: function (id, cb) {
            $http({
                method: "GET",
                url: "/api/thing/" + id
            }).success(cb);
        },
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
        },
        update: function(id, t, cb) {
            $http({
                method: "PUT",
                url: "/api/thing/" + id,
                data: t
            }).success(cb);

        }
    }
}

// Events

// UI

// Rules