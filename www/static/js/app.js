var __APP = angular.module("gohomeapp", [])

// Filters
__APP.filter("htmlcontent", function($sce){
    return function(input) {
        return $sce.trustAsHtml(input);
    }
});

__APP.controller("ctrl-dashboard", function($scope, $http, $interval) {
    $scope.model = {}

    ReloadDashboard($http, $scope);
    $interval(function () {
        ReloadDashboard($http, $scope);
    }, 10000, 0)
});

function ReloadDashboard($http, $scope) {
    var resp = $http.get("/api/ui/dashboard");
    resp.success(function(data, status, headers, config) {
        $scope.model = data;
    });

}
