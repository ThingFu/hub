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
        },
        count: function (id, cb) {
            $http({
                method: "GET",
                url: "/api/thing/" + id + "/event/count"
            }).success(cb);
        },
        action: function (id, action, data, cb) {
            $http({
                method: "POST",
                url: "/api/thing/" + id + "/action/" + action,
                data: data
            }).success(cb);
        }
    }
}

// Events

// UI

// Rules
function RulesManager($http) {
    return {
        save: function (file, id, content, cb) {
            $http({
                method: "POST",
                url: "/api/rule/" + id,
                data: {
                    file: file,
                    content: content
                }
            }).success(cb);
        },
        get: function (id, cb) {
            $http({
                method: "GET",
                url: "/api/rule/" + id
            }).success(cb);
        },
        getAll: function (cb) {
            $http({
                method: "GET",
                url: "/api/rules"
            }).success(cb);
        },
        add: function (file, content, cb) {
            $http({
                method: "PUT",
                url: "/api/rule/" + id,
                data: {
                    file: file,
                    content: content
                }
            }).success(cb);
        },
        delete: function (id, cb) {
            $http({
                method: "DELETE",
                url: "/api/rule/" + id
            }).success(cb);
        }
    }
}
