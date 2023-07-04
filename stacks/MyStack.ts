import { StackContext, Api, EventBus } from "sst/constructs";

export function API({ stack }: StackContext) {
  const api = new Api(stack, "api", {
    routes: {
      "GET /": "packages/functions/src/lambda.handler",
      "GET /{go+}": {
        function: {
          handler: "api/main.go",
          runtime: "go1.x",
        },
      },
    },
  });

  stack.addOutputs({
    ApiEndpoint: api.url,
  });
}
