import * as React from 'react'
import {compose, withHandlers, withState, defaultProps} from 'recompose'

const enhance = compose(
  defaultProps({
    name: 'Test',
  })
)

export const StatelessHome = ({name}) => (
  <div>
    <h1>HOME {name}</h1>
  </div>
)

export const Home = enhance(StatelessHome)
