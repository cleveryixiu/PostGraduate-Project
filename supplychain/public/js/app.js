'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function ($scope, appFactory, $http) {
	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();

	$scope.querySource = function () {
		var id = $scope.query_id;
		appFactory.querySource(id, function (data) {
			$scope.query_source = data;
			if ($scope.query_tuna == "Could not locate tuna") {
				console.log()
				$("#error_query").show();
			} else {
				$("#error_query").hide();
			}
		});
	}
	$scope.queryPart = function () {
		var id = $scope.query_id;
		appFactory.queryPart(id, function (data) {
			$scope.part_source = data;
			if ($scope.query_tuna == "Could not locate tuna") {
				console.log()
				$("#error_query").show();
			} else {
				$("#error_query").hide();
			}
		});
	}
	$scope.queryTransit = function () {
		var id = $scope.query_id;

		// function a (id, callback) {
		// 	.success(function (output) {
		// 		callback(output)
		// 	})
		// }

		appFactory.queryTransit(id, function (data) {
			$scope.transit_source = data;
			if ($scope.query_tuna == "Could not locate tuna") {
				console.log()
				$("#error_query").show();
			} else {
				$("#error_query").hide();
			}
		});
	}
	$scope.queryBlock = function () {
		var id = $scope.query_id;
		appFactory.queryBlock(id, function (data) {
			$scope.id = id;
			$scope.result = data;
			// alert("--------------------------")
			// alert("data:  " + data)
		});
	}

	$scope.queryAllBlock = function () {
		var result = [];

		$http.get('/block/1').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[0] = output;
			}	
		});

		$http.get('/block/2').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[1] = output;
			}
		});

		$http.get('/block/3').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[2] = output;
			}
		});

		$http.get('/block/4').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[3] = output;
			}
		});

		$http.get('/block/5').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[4] = output;
			}
		});

		$http.get('/block/6').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[5] = output;
			}
		});

		$http.get('/block/7').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[6] = output;
			}
		});

		$http.get('/block/8').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[7] = output;
			}
		});

		$http.get('/block/9').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[8] = output;
			}
		});

		$http.get('/block/10').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[9] = output;
			}
		});

		$http.get('/block/11').success(function (output) {
			// alert("result: " + output)
			if(output != "Could not find block") {
				result[10] = output;
			}
		});

		// alert("hello........")

		$scope.block_source = result;
		// alert("ids: " + ids)
		// alert("")
	}

	$scope.queryUser = function () {
		var id = $scope.query_id;
		var usrname = $scope.usn;
		var psw = $scope.psw;
		var uid = hex_hmac_md5(usrname, psw);
		if (uid != id) {
			// alert("uid:"+uid);
			alert("The password or username is not correct");
			window.location.href("/login");
		} else {
			appFactory.queryUser(id, function (data) {
				$scope.part_source = data;
				if ($scope.query_tuna == "Could not locate tuna") {
					console.log()
					$("#error_query").show();
				} else {
					$("#error_query").hide();
					window.location.href = "/index";
				}
			});
		}

	}

	$scope.$watch('$viewContentLoaded', function() {
		$scope.queryAllBlock();
	});
});

app.factory('appFactory', function ($http) {
	var factory = {};
	factory.querySource = function (id, callback) {
		$http.get('/source/' + id).success(function (output) {
			callback(output)
		});
	}
	factory.queryPart = function (id, callback) {
		$http.get('/part/' + id).success(function (output) {
			callback(output)
		});
	}
	factory.queryTransit = function (id, callback) {
		$http.get('/transit/' + id).success(function (output) {
			callback(output)
		});
	}
	factory.queryBlock = function (id, callback) {
		$http.get('/block/' + id).success(function (output) {
			callback(output)
		});
	}

	factory.queryAllBlock = function (callback) {
		var id = 2
		$http.get('/block/' + id).success(function (output) {
			callback(output)
		});
	}

	factory.queryUser = function (id, callback) {
		$http.get('/user/' + id).success(function (output) {
			callback(output)
		});
	}

	factory.getDate = function (callback) {

	}
	return factory;
});
