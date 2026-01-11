import http from 'k6/http';
import { check } from 'k6';

export let options = {
    scenarios: {
        ramping_load: {
            executor: 'ramping-arrival-rate',
            startRate: 1000,
            timeUnit: '1s',
            preAllocatedVUs: 100,
            maxVUs: 200,
            stages: [
                { target: 10000, duration: '1m' },
                { target: 50000, duration: '1m' },
                { target: 0, duration: '30s' },
            ],
        },
    },
};

export default function () {
    let res = http.get('http://localhost');

    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}