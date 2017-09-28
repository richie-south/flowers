import * as React from 'react'
import {
  compose,
  withHandlers,
  withState,
  defaultProps,
  lifecycle
} from 'recompose'
import {withSpinner} from 'react-with-spinner'
import {GET_FLOWER_URL, WATER_FLOWER_URL} from '../../../config/urls'
import {injectIntl} from 'react-intl'
import {ListFlowers} from '../../ui/list-flowers'
import Avatar from 'react-toolbox/lib/avatar/Avatar'
import IconButton from 'react-toolbox/lib/button/IconButton'
import Button from 'react-toolbox/lib/button/Button'
import AlertIcon from 'mdi-react/WaterPumpIcon'
import {post} from '../../../lib/http'
import {Column, Row} from 'styled-material/dist/src/layout'

const enhance = compose(
  withState('data', 'setData', {loading: true}),
  withState('flower', 'setFlower', {}),
  withHandlers({
    waterFlower: ({setFlower, match}) => (amount = 'medium') => {
      post(WATER_FLOWER_URL(match.params.id), {amount})
        .then(response => response.json())
        .then(flower => setFlower(flower))
        .catch(error => console.log(error))
    },
    setLoading: ({setData}) => isLoading =>
      setData(data => ({
        ...data,
        loading: isLoading
      }))
  }),
  lifecycle({
    componentDidMount() {
      fetch(GET_FLOWER_URL(this.props.match.params.id))
        .then(response => response.json())
        .then(flower => {
          this.props.setFlower(flower)
          this.props.setLoading(false)
        })
        .catch(error => console.log(error))
    }
  }),
  withSpinner(),
  injectIntl
)

export const StatelessFlower = ({flower, waterFlower, intl}) => (
  <div>
    <h1>Flower</h1>
    {console.log(flower)}
    <Row horizontal="space-evenly">
      <Row>{flower.flowerType}</Row>
      <Row>{flower.waterIntervall.optimalText}</Row>
      <Row>{flower.waterIntervall.currentText}</Row>
    </Row>
    <Button
      onClick={() => waterFlower('medium')}
      style={{
        position: 'fixed',
        bottom: 16,
        right: 16
      }}
      floating
      accent
      icon={<AlertIcon />}
    />
  </div>
)

export const Flower = enhance(StatelessFlower)
