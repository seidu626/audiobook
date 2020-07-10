import { Injectable } from '@angular/core';
import {
  LanguageServiceClient
} from '../../../proto/service/language/proto/language/language_service_pb_service';
import { environment } from '../../../../environments/environment';
import { LoggerService } from '../../../@core/utils';
import {
  GetRequest, GetResponse,
  ListRequest, ListResponse
} from '../../../proto/service/language/proto/language/language_service_pb';

@Injectable({
  providedIn: 'root',
})
export class LanguageService {
  client: LanguageServiceClient;

  constructor(
    private logger: LoggerService) {
    this.client = new LanguageServiceClient(environment.baseApiEndpoint);
  }

  get(path, val): Promise<object> {
    return new Promise((resolve, reject) => {
      this.logger.log('get.by', path, val);
      const req = new GetRequest();
      req.setId(val);
      this.client.get(req, null, (err, response: GetResponse) => {
        this.logger.log('LanguageService.get.response', response.toObject());
        if (err) {
          return reject(err);
        }
        resolve(response.toObject());
      });
    });
  }

  list(val): Promise<object> {
    return new Promise((resolve, reject) => {
      const req = new ListRequest();
      this.client.list(req, null, (err, response: ListResponse) => {
        if (err) {
          return reject(err);
        }
        resolve(response.toObject());
      });
    });
  }


}
