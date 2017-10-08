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
import {ListFlowers} from '../../ui/list-flowers'
import Avatar from 'react-toolbox/lib/avatar/Avatar'
import IconButton from 'react-toolbox/lib/button/IconButton'
import Button from 'react-toolbox/lib/button/Button'
import WaterPumpIcon from 'mdi-react/WaterPumpIcon'
import {post} from '../../../lib/http'
import {Column, Row} from 'styled-material/dist/src/layout'
import styled from 'styled-components'
import {
  Caption,
  Subhead,
  Title,
  Display4
} from 'styled-material/dist/src/typography'
import {materialColors} from 'styled-material/dist/src/colors'
import {TimelineBox} from '../../ui/timeline'
import {CountDownBlock} from '../../ui/count-down'

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
  withSpinner()
)

const StyledCaption = styled(Caption)`
  color: ${materialColors['grey-600']};
  margin-top: 6px;
`

const StyledRow = styled(Row)`
  padding-top: 16px;
  padding-bottom: 16px;
`

const StyledBlockSubhead = styled(Subhead)`
  font-size: 18px;
  color: ${materialColors['amber-500']};
`

const Block = ({title, data = 'No data!'}) => (
  <StyledRow>
    <Column horizontal="center">
      <Row>
        <StyledBlockSubhead>
          {data.length === 0 ? 'No data!' : data}
        </StyledBlockSubhead>
      </Row>
      <Row>
        <StyledCaption>{title}</StyledCaption>
      </Row>
    </Column>
  </StyledRow>
)

const StyledDivWraper = styled.div`
  background-color: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  border-bottom: 16px;
  margin-bottom: 8px;
`

const Top = ({flower}) => (
  <StyledDivWraper>
    <Column
      horizontal="center"
      style={{
        marginBottom: 24,
        marginTop: 24
      }}
    >
      <Row>
        <Title style={{opacity: 0.8}}>{flower.name}</Title>
      </Row>
      <Row>
        <Subhead style={{opacity: 0.8}}>{flower.flowerType}</Subhead>
      </Row>
    </Column>

    <Row horizontal="space-evenly">
      <Block title="Times watered" data={flower.waterTimeline.length} />
      <Block
        title="Optimal intervall"
        data={flower.waterIntervall.optimal.text}
      />
      <Block
        title="Current intervall"
        data={flower.waterIntervall.current.text}
      />
    </Row>
  </StyledDivWraper>
)

const StyledDisplay4 = styled(Display4)`
  font-weight: 900;
  color: ${materialColors['grey-100']};
  text-transform: uppercase;
  margin-left: 16px;
  overflow: hidden;
  max-width: 86vw;
`

const Top2 = ({flower}) => (
  <div>
    <Column horizontal="flex-start">
      <Row>
        <StyledDisplay4>{flower.name}</StyledDisplay4>
      </Row>
    </Column>
    <Row horizontal="space-evenly">
      <Block title="Times watered" data={flower.waterTimeline.length} />
      <Block
        title="Optimal intervall"
        data={flower.waterIntervall.optimal.text}
      />
      <Block
        title="Current intervall"
        data={flower.waterIntervall.current.text}
      />
    </Row>
  </div>
)

export const StatelessFlower = ({flower, waterFlower}) => (
  <div>
    {console.log(flower)}
    <Column>
      <Top2 flower={flower} />
      <Column horizontal="center">
        <CountDownBlock end={new Date(flower.nextWateringSession)} />
        <TimelineBox timeline={flower.waterTimeline} />
      </Column>
    </Column>
    <Button
      onClick={() => waterFlower('medium')}
      style={{
        position: 'fixed',
        bottom: 16,
        right: 16
      }}
      floating
      accent
      icon={
        <WaterPumpIcon
          style={{
            fill: 'white'
          }}
        />
      }
    />
  </div>
)

export const Flower = enhance(StatelessFlower)
