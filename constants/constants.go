package constants

import "time"

const API_KEY = "RGAPI-641a0b82-0fbe-4c09-a6e0-bc756cf92521"
const URL_CLIENT_FORMAT = "https://127.0.0.1:%s/" // Port number

const CHECK_IF_CLIENT_OPEN_TIME = 10 * time.Second
const CHECK_IF_GAME_STARTED_TIME = 5 * time.Second
const CHECK_IF_GAME_OVER_TIME = 10 * time.Second

const MAC_SETTINGS_DB_PATH = "~/Documents/KisaData"
const WINDOWS_SETTINGS_DB_PATH = "C:/Users/%s/Documents/KisaData"

const LOG_PERIODIC_ON = true
const LOG_DEBUG_ON = true
const LOG_INFO_ON = true
const LOG_ERROR_ON = true
const LOG_FATAL_ON = true