# pow-rplace

A Reddit's r/place-like application using proof of work.

## Description

Inspired by Reddit's r/place, this project involves a pixel map where users can claim individual pixels to draw various designs. However, in this version, a proof-of-work mechanism is used to "protect" each pixel. This means that in order to claim a pixel, users must solve a computational problem. The difficulty of the problem is dynamically adjusted based on the number of users attempting to update the same pixel.

Each pixel has a state and a queue of incoming actions. Every minute, the queue is processed, and the pixel’s state is updated. Additionally, the difficulty (or "bits") required for claiming the pixel is recalculated.

The difficulty curve should be exponential enough to prevent a large source of computational power from monopolizing a pixel for an extended period (i.e., filling up the queue).
TODO:

    Build a map/queue system to store actions.
    Expose an API to get the expected difficulty (bits) for a given pixel.
    Expose an API to post a new action.
    Expose an API to get the current state of the map or a specific pixel.

## Architecture
### Engine

The engine is responsible for maintaining the map and its action queues. It must be highly performant and secure. The state should be stored both in-memory (for real-time processing) and on disk (to keep a history of changes). This is the only component that should handle storage in the backend.
Backend API

### Backend API
The backend API must expose three endpoints:

    GET bits – Retrieve the expected difficulty for a given pixel.
    POST action – Submit a new action to claim a pixel.
    GET state – Retrieve the current state of either the entire map or a specific pixel.