import React from 'react'
import Route from 'react-router/Route'
import Switch from 'react-router/Switch'
import {Column} from 'styled-material/dist/src/layout'
import {Home} from './pages/home/home'

export const TopRouter = () => (
  <Column>
    <Switch>
      <Route path="/" component={Home} />
    </Switch>
  </Column>
)
