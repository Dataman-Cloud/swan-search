/**
 * Created by my9074 on 2017/1/9.
 */
export class SearchController {
  constructor(searchBackend, moment, BACKEND_URL_BASE) {
    'ngInject';

    this.searchBackend = searchBackend;
    this.monitorBase = BACKEND_URL_BASE.monitorBase;
    this.swanBase = BACKEND_URL_BASE.swanBase;
    this.appDomain = BACKEND_URL_BASE.appDomain;
    this.keyword = '';
    this.clusters = [];
    this.apps = [];
    this.end = moment().unix();
    this.start = moment().subtract(120, 'minutes').unix();
    this.activate();
  }

  activate() {

  }

  searchClusters() {
    this.apps = [];

    this.searchBackend.searchApps(this.keyword).get(data => {
      if (Array.isArray(data.data)) {
        this.apps = data.data.filter(item => item.Type === 'app').map(app => {
          app.formatID = app.ID.split('-').join('.');
          return app;
        });
      }
    })
  }
}
