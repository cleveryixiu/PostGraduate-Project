
'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	$scope.querySource = function(){
		var id = $scope.query_id;
		appFactory.querySource(id, function(data){
			$scope.query_source = data;
			if ($scope.query_tuna == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}
	$scope.queryPart = function(){
		var id = $scope.query_id;
		appFactory.queryPart(id, function(data){
			$scope.part_source = data;
			if ($scope.query_tuna == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}
	$scope.queryTransit = function(){
		var id = $scope.query_id;
		appFactory.queryTransit(id, function(data){
			$scope.transit_source = data;
			if ($scope.query_tuna == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}
    $scope.queryBlock = function(){
        var id = $scope.query_id;
        appFactory.queryBlock(id, function(data){

            if ($scope.query_tuna == "Could not locate tuna"){
                console.log()
                $("#error_query").show();
            } else{
                $("#error_query").hide();
                $scope.id = id;
                $scope.result = data;
            }
        });
    }
    $scope.queryUser = function(){
        var id = $scope.query_id;
        var usrname =$scope.usn;
        var psw = $scope.psw;
        var uid = hex_hmac_md5(usrname,psw);
        if(uid!=id){
        	alert("uid:"+uid);
            alert("您输入的用户名密码与用户id不匹配，请重新输入");
            window.location.href("/login");
		}else{
            appFactory.queryUser(id, function(data){
                $scope.part_source = data;
                if ($scope.query_tuna == "Could not locate tuna"){
                    console.log()
                    $("#error_query").show();
                } else{
                    $("#error_query").hide();
                    window.location.href="/";
                }
            });
		}

    }
});

app.factory('appFactory', function($http){
	var factory = {};
	factory.querySource = function(id, callback){
    	$http.get('/source/'+id).success(function(output){
			callback(output)
		});
	}
	factory.queryPart = function(id, callback){
    	$http.get('/part/'+id).success(function(output){
			callback(output)
		});
	}
	factory.queryTransit = function(id, callback){
    	$http.get('/transit/'+id).success(function(output){
			callback(output)
		});
	}
    factory.queryBlock = function(id, callback){
        $http.get('/block/'+id).success(function(output){
            callback(output)
        });
    }
    factory.queryUser = function(id, callback){
        $http.get('/user/'+id).success(function(output){
            callback(output)
        });
    }
	return factory;
});
