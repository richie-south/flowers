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

const enhance = compose(injectIntl)

export const StatelessListFlowers = ({flowers, intl}) => (
  <List>
    {flowers.map(({id, name, nextWateringSession}) => (
      <ListItem
        key={id}
        caption={name}
        legend={'KING'}
        leftIcon="local_florist"
      />
    ))}
  </List>
)

export const ListFlowers = enhance(StatelessListFlowers)
