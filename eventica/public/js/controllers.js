'use strict';

/* Controllers */

angular.module('mango.controllers', [])
    .controller('UserController', ['$scope', '$http', 'UserService', '$rootScope', '$location', function($scope, $http, UserService, $rootScope, $location) {
        $scope.user = {
            name: "",
            lastname: "",
            email: "",
            email2: "",
            password: "",
            rememberme: true
        };

        $scope.registerUser = function(){
            $http.post('/api/user/register', $scope.user)
                .success(function(data){
                    $scope.user = data.user;
                    $rootScope.flashes = data.flashes;
                })
        };

        $scope.loginUser = function(){
            $http.post('/api/user/authenticate', $scope.user)
                .success(function(data,status,headers,config){
                    if (data["valid"] === true) {
                        $rootScope.$broadcast("userAuthstateChanged", data.user);
                        $location.path(data["redirect"]);
                    } else {
                        $rootScope.flashes = data.flashes;
                        $scope.user.password = "";
                    }
                })
        };
    }])
    .controller('UserResetPasswordRequestController', ['$scope', '$http', '$rootScope', function($scope, $http, $rootScope) {
        $scope.user = {
            email: ""
        };
        $scope.sendRequest = function(){
            $http.post('/api/user/reset_request', $scope.user)
                .success(function(data,status,headers,config){
                    $rootScope.flashes = data.flashes;
                })
        };
    }])
    .controller('UserResetPasswordController', ['$scope', '$http', '$location', '$rootScope', function($scope, $http, $location ,$rootScope) {
        $scope.token = $location.search();

        $scope.resetPassword = function(){
            $http.post('/api/user/reset_password', $scope.token)
                .success(function(data,status,headers,config){
                    $rootScope.flashes = data.flashes;
                })
        }

    }])
    .controller('UserLogoutController', ['$scope', '$http', 'UserService', '$rootScope', function($scope, $http, UserService, $rootScope) {
        $scope.$on('$routeChangeSuccess', function () {
            $http.post('/api/user/endsession')
                .success(function(data){
                    $rootScope.flashes = data.flashes;
                    $rootScope.$broadcast("userAuthstateChanged",{name:null});
                })
        });

    }])
    .controller('UserProfileController', ['$scope', '$http', 'UserService', '$rootScope', '$location', function($scope, $http, UserService, $rootScope, $location) {
        $scope.$on('$routeChangeSuccess', function () {
            $http.get('/api/user/profile', UserService)
                .success(function(data,status,headers,config){
                    if (data["redirect"] != null){
                        $location.path(data["redirect"]);
                    } else {
                        $scope.profile = data.profile;
                    }
                })
        });
        $scope.updateProfile = function(){
            $http.post('/api/user/profile', $scope.profile)
                .success(function(data){
                    $rootScope.flashes = data.flashes;
            });
        }
    }])
;
