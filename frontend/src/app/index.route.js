export function routerConfig ($stateProvider, $urlRouterProvider) {
  'ngInject';
  $stateProvider
    .state('home', {
      templateUrl: 'app/main/main.html',
      controller: 'MainController',
      controllerAs: 'main'
    })
    .state('home.search', {
    url: '/search',
    templateUrl: 'app/main/search/search.html',
    controller: 'SearchController',
    controllerAs: 'vm'
  });

  $urlRouterProvider.otherwise($injector => {
    let $state = $injector.get('$state');
    $state.go('home.search');
  });
}
