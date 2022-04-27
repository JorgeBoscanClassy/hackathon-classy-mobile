import { StyleSheet } from 'react-native';
import { useEffect , useState } from 'react'
import { Text, View } from '../components/Themed';
import { RootTabScreenProps } from '../types';
import AnimatedNumbers from 'react-native-animated-numbers';


export default function Home({ navigation }: RootTabScreenProps<'TabOne'>) {

  const [ raised, setRaised ] = useState(0);

  const amount_raised = 1560300;
  

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
      <Text style={styles.sectionHeader}>Donations</Text>


      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
    </View>
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
