# Real-Time Hyperlocal Feed and Notification Engine

Portfolio-ready Go + Node.js implementation using Redis Streams and MongoDB, with k3s manifests, Prometheus alerts and k6 load testing.

Run Redis and MongoDB, then `go run ./ingestion`; in `gateway`, run `npm install && npm start`; finally `k6 run loadtest.js`.
