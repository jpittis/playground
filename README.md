What I use to get hacking on a new idea.

- app/ is an example Go app which produces prometheus metrics
- Grafana is running on localhost:3000 and is backed by prometheus
- cadvisor monitors container resource usage

Just docker-compose up, visit localhost:3000 in your browser, and get started!

Warning: cadvisor setup doesn't seem to work on macOS
