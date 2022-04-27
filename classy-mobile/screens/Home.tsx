import { StyleSheet } from 'react-native';
import { useEffect , useState } from 'react'
import { Text, View } from '../components/Themed';
import { RootTabScreenProps } from '../types';
import AnimatedNumbers from 'react-native-animated-numbers';
import { Highlights } from '../components/Highlights';
import { ScrollView } from 'react-native-gesture-handler';

export default function Home({ navigation }: RootTabScreenProps<'TabOne'>) {

  const [ raised, setRaised ] = useState(0);

  const amount_raised = 1560300;
  const highlights = [{title : "$15,603.00", label:"Average Transaction Size" },
                      {title : "$5,211,197.00", label:"Total Transactions" },
                      {title : "5,632", label:"Active Campaigns" },
                      {title : "$131,632", label:"Average Raised" },
                      {title : "14", label:"Average Transactions" },]
  

  useEffect(() => {
    
    //Initial animation and set the number
    setTimeout(()=> {   setRaised(amount_raised) }, 100);
    
    //Random donations every 2 seconds
    setInterval(() => {

      let max = 500;
      let min = 50000;
      setRaised(Math.floor(Math.random() * (max - min)) + amount_raised);

    },4000);


  }, [])


  return (
    <ScrollView>
    <View style={styles.container}>
      <View style={{ flexDirection:'row'}}><Text style={styles.title}>$</Text><AnimatedNumbers
        includeComma
        animateToNumber={raised}
        fontStyle={styles.title}
      /></View>
      <Text style={styles.subTitle}>Rasied this week</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <View style={{ marginLeft:-20}}>
      </View>
      <Text style={styles.sectionHeader}>Highlights</Text>
      <Highlights data={highlights} />
      <Text style={styles.sectionHeader}>Donations</Text>


      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
    </View>
    </ScrollView>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
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
    marginTop:10,
    fontWeight:'bold'
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
  }
});
