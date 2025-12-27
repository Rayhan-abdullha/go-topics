import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
  vus: 1000,          // virtual users
  iterations: 10000, // total requests
};

export default function () {
  http.get('https://your-api-url.com/api/test');
  sleep(1);
}
