import { StyleSheet } from 'react-native';
import { useEffect , useState } from 'react'
import { Text, View } from '../components/Themed';
import { RootTabScreenProps } from '../types';
import AnimatedNumbers from 'react-native-animated-numbers';
import { Highlights } from '../components/Highlights';
import { ListItems } from '../components/ListItems';
import { ScrollView } from 'react-native-gesture-handler';
import { VictoryLine, VictoryChart, VictoryAxis } from "victory-native";


export default function Home({ navigation }: RootTabScreenProps<'TabOne'>) {

  const [ raised, setRaised ] = useState(0);
  const [ hour, setHour ] = useState(0);

  const amount_raised = 1560300;
  const highlights = [{title : "$15,603.00", label:"Average Transaction Size" },
                      {title : "$5,211,197.00", label:"Total Transactions" },
                      {title : "5,632", label:"Active Campaigns" },
                      {title : "$131,632", label:"Average Raised" },
                      {title : "14", label:"Average Transactions" },]

const donations = [{title : "Omid Borjian", label:"Average Transaction Size", right: "$530" },
                      {title : "Tammen Bruccoleri", label:"Total Transactions", right: "$420"  },
                      {title : "Emad Borjian", label:"Active Campaigns", right: "$380"  },
                      {title : "Chris Himes", label:"Average Raised" , right: "$850" }]
  
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


  const getHour = () => {
    const date = new Date();
    const hour = date.getHours()
    setHour(hour);
   }

  useEffect(() => {
    
    //Initial animation and set the number
    setTimeout(()=> {   setRaised(amount_raised) }, 100);

    //Get the time of the day for a greeting
    getHour();
    
    //Random donations every 2 seconds
    setInterval(() => {

      let max = 500;
      let min = 50000;
      setRaised(Math.floor(Math.random() * (max - min)) + amount_raised);

    },4000);


  }, [])


  return (
    <ScrollView style={{ backgroundColor : '#fff'}}>
    <View style={styles.container}>
    <View style={styles.spacer}></View>
    <Text style={styles.greeting}>{hour < 12 ? "Good Morning" : "Good evening"}, Omid!</Text>
      <View style={{ flexDirection:'row'}}><Text style={styles.title}>$</Text><AnimatedNumbers
        includeComma
        animateToNumber={raised}
        fontStyle={styles.title}
      /></View>
      <Text style={styles.subTitle}>Rasied this week</Text>
  
      <View style={{ marginLeft:-20}}>
      </View>
      <View style={{ marginLeft:-20}}>
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
    axis: {stroke: "transparent"}, 
    ticks: {stroke: "transparent"},
    tickLabels: { fill:"transparent"} 
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
    backgroundColor:'#fff',
    padding:20,
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
  },
  title: {
    fontSize: 40,
    fontWeight: 'bold',
  },
  subTitle : {
    fontSize:15,
    marginTop:10
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
  sectionHeader:{
    fontSize:18,
    fontWeight:'bold',
    marginTop:10
  },
  appBackground:{
    backgroundColor:'#fff'
  },
  spacer:{
    marginBottom:25
  },
  greeting: {
    fontSize:30,
    marginBottom:20
  }
});
