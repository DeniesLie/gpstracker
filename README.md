# gpstracker
Golang labs project

## Description
Simple http service that provides functionality to create routes, send coordinates of "current" position (waypoints), and get info about traveled route (total distance, duration, average speed, etc.)

## Endpoints
### Routes
- **GET**: */routes* <br>
  <ins>Description</ins>: Get all routes<br>
  <ins>Response body</ins>: RouteGetDto[]<br>
  <ins>Status codes</ins>: <br>
    + 200 Success <br>
    
 - **GET**: */routes/info* <br>
  <ins>Description</ins>: Get info about route<br>
  <ins>Query params</ins>: routeId:int, totalDistance, averageSpeed  <br>
  <ins>Response body</ins>: RouteInfoDto<br>
  <ins>Status codes</ins>: <br>
    + 200 Success <br>
    + 404 Route with id {routeId} was not found <br>
    
- **POST**: */routes* <br>
  <ins>Description</ins>: Create new route <br>
  <ins>Request body</ins>: RoutePostDto<br>
  <ins>Response body</ins>: RouteGetDto<br>
  <ins>Status codes</ins>: <br>
    + 201 Created <br>
    + 400 Route with same name already exists <br>
    
 - **POST**: */routes/complete*<br>
  <ins>Description</ins>: Complete the route, so new waypoints can't be added. The request is idempotent, it will return success if route is already completed<br>
  <ins>Query params</ins>: routeId:int <br>
  <ins>Status codes</ins>: <br>
    + 200 Success <br>
    + 404 Route with id {routeId} was not found <br>
    
- **DELETE**: */routes* <br>
  <ins>Description</ins>: Delete the route with all related waypoints <br>
  <ins>Query params</ins>: routeId:int <br>
  <ins>Status codes</ins>: <br>
    + 200 Success <br>
    + 404 Route with id {routeId} was not found <br>
    
### Waypoints
  - **GET**: */waypoints* <br>
  <ins>Description</ins>: Get waypoints for a route in specified time interval <br>
  <ins>Query params</ins>: routeId:int, after:string, before:string  <br>
  <ins>Status codes</ins>: <br>
    + 200 Success <br>
    + 400 after/before times values are not valid <br>
    + 404 Route with id {WaipointPostDto.routeId} was not found <br>

  - **POST**: */waypoints* <br>
  <ins>Description</ins>: Add waypoint to a route <br>
  <ins>Request body</ins>: WaipointPostDto <br>
  <ins>Status codes</ins>: <br>
    + 201 Created <br>
    + 400 Coordinates values are out of range <br>
    + 404 Route with id {WaipointPostDto.routeId} was not found <br>
    + 409 Can't add waypoint when route is in Completed state <br>
    
  - **POST**: */waypoints/addRange* <br>
  <ins>Description</ins>: Add timed waypoints to a route <br>
  <ins>Request body</ins>: TimedWaypointsPostDto <br>
  <ins>Status codes</ins>: <br>
    + 201 Created <br>
    + 400 Coordinates values are out of range <br>
    + 404 Route with id {WaipointPostDto.routeId} was not found <br>
    + 409 Can't add waypoint when route is in Completed state <br>
    + 409 Waypoint with same time already exists in the route <br>
    + 409 Adding a waypoint with time before latest waypoint is not allowed<br>

## Dtos
- RoutePostDto
- RouteGetDto
- RouteInfoDto
- WaypointGetDto
- WaypointPostDto
- TimedWaypointsPostDto

## Data model
![image](https://user-images.githubusercontent.com/59698344/217930179-144b1649-307b-4212-9a71-88e44863f8e4.png)
![image](https://user-images.githubusercontent.com/59698344/217930254-5eddd87e-176c-431b-8135-6c76e12e2a6e.png)

