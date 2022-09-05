import {HelloWorldDTO} from "./types";

export const getHelloWorld: () => Promise<HelloWorldDTO> =
  () => fetch('/api/hello')
    .then(response => response.json())
