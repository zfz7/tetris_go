import React from 'react';
import {cleanup, render} from '@testing-library/react';
import {HelloWorld} from "./HelloWorld";
import {getHelloWorld} from "../exchange/GetHelloWorld";
import {act} from "react-dom/test-utils";


jest.mock('../exchange/GetHelloWorld')
const getHelloWorldMock = getHelloWorld as jest.MockedFunction<typeof getHelloWorld>


describe('HelloWorld', () => {
  beforeEach(async () => {
    getHelloWorldMock.mockResolvedValue({message: 'Hi'})
  })
  afterEach(() => {
    cleanup()
    getHelloWorldMock.mockReset()
  })
  it('it displays hello world', async () => {
    const {container} = render(<HelloWorld/>)
    await act(async () => {
      await getHelloWorldMock
    })
    expect(container.innerHTML).toContain("Hello World: Hi")
  })
})