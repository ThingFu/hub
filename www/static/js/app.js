var __APP = angular.module("gohomeapp", []);
var ThingService = function($http) {
    this.get = function (id, cb) {
        $http.get("/api/thing/" + id).success(cb);
    }
};
ThingService.$inject = ["$http"];

__APP.service("$things", function($http){
    return new ThingService($http)
});
