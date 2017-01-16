export class searchBackend {
  constructor($resource, BACKEND_URL_BASE) {
    'ngInject';

    this.$resource = $resource;
    this.BACKEND_URL_BASE = BACKEND_URL_BASE;
  }

  searchApps(data) {
    return this.$resource(`${this.BACKEND_URL_BASE.defaultBase}/search/v1/luckysearch`, {keyword: data});
  }
}
