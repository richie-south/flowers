import * as React from 'react'
import {
  compose,
  withHandlers,
  withState,
  defaultProps,
  lifecycle
} from 'recompose'
import {withSpinner} from 'react-with-spinner'
import List from 'react-toolbox/lib/list/List'
import ListItem from 'react-toolbox/lib/list/ListItem'
import {GET_FLOWER_LIST_URL} from '../../config/urls'
import {injectIntl} from 'react-intl'
import {CountDownSpan} from './count-down'
import Redirect from 'react-router/Redirect'

const enhance = compose(withState('redirect', 'setRedirect', ''), injectIntl)

export const StatelessListFlowers = ({
  redirect,
  setRedirect,
  viewFlower,
  flowers,
  intl
}) =>
  redirect ? (
    <Redirect to={redirect} />
  ) : (
    <List>
      {flowers.map(({id, name, nextWateringSession}) => (
        <ListItem
          key={id}
          onClick={() => setRedirect(id)}
          caption={name}
          legend={<CountDownSpan end={new Date(nextWateringSession)} />}
          leftIcon="local_florist"
        />
      ))}
    </List>
  )

export const ListFlowers = enhance(StatelessListFlowers)
