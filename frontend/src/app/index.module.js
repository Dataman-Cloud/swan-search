/* global malarkey:false, moment:false */

import { config } from './index.config';
import { routerConfig } from './index.route';
import { runBlock } from './index.run';
import { MainController } from './main/main.controller';
import { SearchController } from './main/search/search.controller';
import { httpInterceptor } from '../app/utils/service/httpInterceptor.service';
import { searchBackend } from '../app/main/search/service/search-backend.service';

angular.module('frontend', ['ngAnimate', 'ngCookies', 'ngSanitize', 'ngMessages', 'ngAria', 'ngResource', 'ui.router', 'ngMaterial',
  'ui-notification', 'angular-loading-bar'])
  .constant('moment', moment)
  .constant('BACKEND_URL_BASE', {
    defaultBase: "http://192.168.1.155:9888",
    swanBase: "http://192.168.1.155:3000",
    monitorBase: "http://192.168.1.75:5098/ui/monitor/chart"
  })
  .config(config)
  .config(routerConfig)
  .run(runBlock)
  .service('httpInterceptor', httpInterceptor)
  .service('searchBackend', searchBackend)
  .controller('MainController', MainController)
  .controller('SearchController', SearchController);
