FROM oven/bun:1 AS base
WORKDIR /usr/src/app

FROM base AS install
RUN mkdir -p /temp
COPY package.json bun.lockb /temp/
RUN cd /temp && bun install --frozen-lockfile --production

FROM install AS prerelease
COPY --from=install /temp/node_modules node_modules
COPY . .
ENV NODE_ENV=production
RUN bun run build

FROM base AS release
COPY --from=install /temp/node_modules node_modules
COPY --from=prerelease /usr/src/app/.output .

USER bun
EXPOSE 3000
ENV NODE_ENV=production 
ENTRYPOINT [ "bun", "run", "server/index.mjs" ]