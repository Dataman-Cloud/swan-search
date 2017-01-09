import { BACKEND_URL_BASE } from '../../../conf'

export class searchBackend {
  constructor($resource) {
    'ngInject';

    this.$resource = $resource;
  }

  cluster() {
    return this.$resource(BACKEND_URL_BASE.defaultBase + '/v1/stats');
  }
}
