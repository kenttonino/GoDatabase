import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  thresholds: {
    // Assert that 99% of requests finish within 3000ms.
    http_req_duration: ["p(99) < 3000"],
  },
  // Ramp the number of virtual users up and down.
  stages: [
    { duration: "30s", target: 15 },
    { duration: "1m", target: 15 },
    { duration: "20s", target: 0 },
  ],
  noConnectionReuse: true,
  noVUConnectionReuse: true,
};

// Simulated user behavior.
export default function () {
  const res = http.get("http://localhost:3000/ready");
  check(res, { "status was 200": (r) => r.status == 200 });
  sleep(1);
}
