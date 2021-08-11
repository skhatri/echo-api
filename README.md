### Echo
Run

```
docker run -p 8080:8080 -it skhatri/echo-api
```

Test

```
#index
curl http://localhost:8080/

#echo request headers and body
curl -H"content-type: application/json" -d '{"status": "all_done", "message": "build-completed"}' http://localhost:8080/echo

#echo request headers
curl http://localhost:8080/echo
```
