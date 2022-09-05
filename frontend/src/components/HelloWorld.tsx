import React, {useEffect, useState} from 'react'
import {getHelloWorld} from "../exchange/GetHelloWorld";
import {HelloWorldDTO} from "../exchange/types";

export const HelloWorld: React.FC = () => {
  const [helloWorld, setHelloWorld] = useState<HelloWorldDTO>({message:"pending..."})
  useEffect(() => {
    getHelloWorld().then(response => setHelloWorld(response))
  }, [])

  return (
    <div>Hello World: {helloWorld.message}</div>
  )
}
