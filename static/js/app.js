function AppCtrl($scope, $http) {
  $scope.lista = [];

  var refresh = function() {
    return $http.get('/tarea/').
      success(function(data) { $scope.lista = data.Tasks; })  
  };

  $scope.add = function() {
  
    $http.post('/tarea/', {Title: $scope.todoText}).     
      success(function() {
        refresh().then(function() {        
          $scope.todoText = '';
        })
      });
  };
 
}