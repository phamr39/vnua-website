const isProductionEnv = process.env.NODE_ENV === "production";

export const API_DOMAIN = isProductionEnv
? "http://vnuapharma.com.vn:3000"
: "http://vnuapharma.com.vn:3000";

export const BASE_URL = isProductionEnv ? "/vnua/" : "/";

export const DOMAIN_TITLE = "vnua";
