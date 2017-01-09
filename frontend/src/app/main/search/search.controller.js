/**
 * Created by my9074 on 2017/1/9.
 */
export class SearchController {
  constructor(searchBackend) {
    'ngInject';

    this.searchBackend = searchBackend;
    this.clusters = [1, 2, 3];
    this.apps = [1, 2, 3, 4, 5];
    this.activate();
  }

  activate() {

  }


  searchClusters() {
    this.clusters = [];
    this.apps = [];

    //TODO GET CLUSTER AJAX
    this.searchBackend.cluster().get(data => {

    })
  }
}
