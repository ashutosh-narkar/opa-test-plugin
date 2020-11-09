import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '3s', target: 6000 },
        { duration: '5s', target: 8000 },
        { duration: '10s', target: 0 },
    ],
    // thresholds: {
    //     http_req_duration: ['p(99)<1500'], // 99% of requests must finish within 1.5s.
    // },
};


export default function () {
    let res = http.get('http://bob:password@192.168.64.20:30556/api/v1/products');

    check(res, { 'Response status was 200': (r) => r.status == 200 });

    check(res, { 'Response status was 503': (r) => r.status == 503 });

    check(res, { 'Response status was 508': (r) => r.status == 508 });

    check(res, { 'Response status was 400': (r) => r.status == 400 });

    sleep(1);
}