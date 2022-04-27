import { StyleSheet } from 'react-native';
import { useEffect, useState } from 'react'
import { Text, View } from '../components/Themed';
import { RootTabScreenProps } from '../types';
import AnimatedNumbers from 'react-native-animated-numbers';
import { Highlights } from '../components/Highlights';
import { ListItems } from '../components/ListItems';
import { ScrollView } from 'react-native-gesture-handler';
import { VictoryLine, VictoryChart, VictoryAxis } from "victory-native";
import EventSource from "react-native-sse";

export default function Home({ navigation }: RootTabScreenProps<'TabOne'>) {
  const es = new EventSource("http://localhost:4000/sse/subscribe");

  const [raised, setRaised] = useState(0);

  const amount_raised = 1560300;
  const highlights = [{ title: "$15,603.00", label: "Average Transaction Size" },
  { title: "$5,211,197.00", label: "Total Transactions" },
  { title: "5,632", label: "Active Campaigns" },
  { title: "$131,632", label: "Average Raised" },
  { title: "14", label: "Average Transactions" },]

  const donations = [{ id: 1, title: "Omid Borjian", label: "Average Transaction Size", right: "$530" },
  { id: 2, title: "Tammen Bruccoleri", label: "Total Transactions", right: "$420" },
  { id: 3, title: "Emad Borjian", label: "Active Campaigns", right: "$380" },
  { id: 4, title: "Chris Himes", label: "Average Raised", right: "$850" }]

  //Charts data
  const thisWeek = [
    { x: 0, y: 0 },
    { x: 2, y: 3 },
    { x: 3, y: 5 },
    { x: 4, y: 4 },
    { x: 5, y: 6 }
  ];

  const lastWeek = [
    { x: 0, y: 0 },
    { x: 2, y: 5 },
    { x: 3, y: 2 },
    { x: 4, y: 3 },
    { x: 5, y: 4 }
  ];

  //{"raisedThisWeek":8408,"donations":[{"name":"Omid Borijan","time":"2022-04-27T16:29:05.05123-07:00","campaign":"WorldCentral","amount":7008},{"name":"Tammen K","time":"2022-04-27T16:29:05.05123-07:00","campaign":"Tunnels to Towers","amount":5528},{"name":"Emad B","time":"2022-04-27T16:29:05.05123-07:00","campaign":"Tunnels to Towers","amount":3653}]}

  interface StatsEvent {
    raisedThisWeek: number;
  }

  useEffect(() => {

    //Initial animation and set the number
    setTimeout(() => { setRaised(amount_raised) }, 100);

    //Random donations every 2 seconds
    // setInterval(() => {

    //   let max = 500;
    //   let min = 50000;
    //   setRaised(Math.floor(Math.random() * (max - min)) + amount_raised);

    // }, 4000);

    es.addEventListener("open", (event) => {
      console.log("Open SSE connection.");
    });

    es.addEventListener("message", (event) => {
      console.log("New message event:", event.data);
      const statsEvent = JSON.parse(event.data) as StatsEvent;
      setRaised(statsEvent.raisedThisWeek)
    });

    es.addEventListener("error", (event) => {
      if (event.type === "error") {
        console.error("Connection error:", event.message);
      } else if (event.type === "exception") {
        console.error("Error:", event.message, event.error);
      }
    });

    es.addEventListener("close", (event) => {
      console.log("Close SSE connection.");
    });

  }, [])

  return (
    <ScrollView style={{ backgroundColor: '#fff' }}>
      <View style={styles.container}>
        <View style={styles.spacer}></View>
        <View style={{ flexDirection: 'row' }}><Text style={styles.title}>$</Text><AnimatedNumbers
          includeComma
          animateToNumber={raised}
          fontStyle={styles.title}
        /></View>
        <Text style={styles.subTitle}>Rasied this week</Text>

        <View style={{ marginLeft: -20 }}>
        </View>
        <View style={{ marginLeft: -20 }}>
          <VictoryChart>
            <VictoryLine
              interpolation="natural"
              style={{ data: { stroke: "#f4775e", strokeWidth: 7, strokeLinecap: "round" } }}
              data={thisWeek}
              animate={{
                duration: 2000,
                onLoad: { duration: 1000 }
              }}
            />
            <VictoryLine
              data={lastWeek}
              style={{ data: { stroke: "#E2E2E2", strokeWidth: 7, strokeLinecap: "round" } }}
              interpolation="natural"
              animate={{
                duration: 2000,
                onLoad: { duration: 2000 }
              }}
            />
            <VictoryAxis style={{
              axis: { stroke: "transparent" },
              ticks: { stroke: "transparent" },
              tickLabels: { fill: "transparent" }
            }} />
          </VictoryChart>
        </View>

        <Text style={styles.sectionHeader}>Highlights</Text>


        <Highlights data={highlights} />

        <View style={styles.spacer}></View>
        <Text style={styles.sectionHeader}>Donations</Text>

        <View style={styles.spacer}></View>
        <ListItems data={donations} />


      </View>
    </ScrollView>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    padding: 20,
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
  },
  title: {
    fontSize: 40,
    fontWeight: 'bold',
  },
  subTitle: {
    fontSize: 15,
    marginTop: 10,
    fontWeight: 'normal'
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
  sectionHeader: {
    fontSize: 18,
    fontWeight: 'bold',
    marginTop: 10
  },
  appBackground: {
    backgroundColor: '#fff'
  },
  spacer: {
    marginBottom: 25
  }
});
