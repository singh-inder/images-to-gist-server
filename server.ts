import "dotenv/config";
import express, { type ErrorRequestHandler } from "express";
import path from "node:path";
import { pipeline } from "node:stream/promises";
import helmet from "helmet";
import parseUrl from "parse-url";
import createHttpError, { isHttpError } from "http-errors";
import got from "got";
import { Base64Decode } from "base64-stream";

import hosts from "./lib/hosts.js";

const app = express();

app.use(helmet({ crossOriginResourcePolicy: { policy: "cross-origin" } }));

const WEEK_IN_SECONDS = 60 * 60 * 24 * 7;

app.get("/", async (req, res, next) => {
  try {
    const url = req.query.url;

    if (typeof url !== "string") return next(createHttpError(400, "No Url provided"));

    const parsedUrl = parseUrl(url, true);

    if (parsedUrl.parse_failed) return next(createHttpError(400, "Invalid url"));

    if (!hosts.includes(parsedUrl.host))
      return next(createHttpError(400, "Invalid url"));

    const mimeType = path.basename(url).split(".").pop() || "jpeg";

    res.setHeader("Content-type", `image/${mimeType}`);
    res.setHeader("Cache-Control", `private, max-age=${WEEK_IN_SECONDS}`);

    await pipeline(
      got.stream(`https://${parsedUrl.host}${parsedUrl.pathname}`),
      new Base64Decode(),
      res
    );
  } catch (error) {
    next(error);
  }
});

app.use(((err, req, res, next) => {
  if (isHttpError(err)) {
    res.status(err.status).json({ message: err.message });
  } else {
    res.sendStatus(500);
  }
}) as ErrorRequestHandler);

const PORT = process.env.PORT || 5000;

app.listen(PORT, () => console.log(`server running on port ${PORT}`));
