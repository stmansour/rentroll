"use strict";

// base url for API endpoints
export const API_BASE_URL = "http://localhost:8270";

// path for AIR Receipt application
export const RECEIPT_APPLICATION_PATH = '/rhome';
// path for AIR Roller application
export const ROLLER_APPLICATION_PATH = '/home';

// API version
export const API_VERSION = "v1";

// Unset business id
// export const BID = -1;
export const BID = 1; //TODO(Akshay): Handle dynamically

// Success flag to match with API response status
export const API_RESPONSE_SUCCESS_FLAG = 'success';

// HTTP STATUS CODE
export const HTTP_OK_STATUS = 200;

// Application cookie's key
export const APPLICATION_COOKIE = 'airoller';

// wait time in application
export const WAIT_TIME = 2000;
export const PAGE_LOAD_TIME = 2000;
export const LOGIN_WAIT_TIME = 2000;

// aliasing for cypress commands