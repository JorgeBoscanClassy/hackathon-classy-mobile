import { StyleSheet , Dimensions, Touchable, Image, Animated } from 'react-native';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { useEffect, useState, useRef } from 'react';
import { ScrollView, TouchableOpacity } from 'react-native-gesture-handler';
import { ListItems } from '../components/ListItems'

const { width } = Dimensions.get('window');


export default function CampaignDetails({ navigation }) {

   
    const progressWidth = 100;

    const data = {

            title:"Give Now to Help Families",
            raised:'$145,667.97',
            cover : 'https://images.unsplash.com/photo-1604176354204-9268737828e4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1160&q=80'


    }

      //Charts data
      const donations = [{ id: 1, title: "Omid Borjian", label: "Average Transaction Size", right: "$530" },
      { id: 2, title: "Tammen Bruccoleri", label: "Total Transactions", right: "$420" },
      { id: 3, title: "Emad Borjian", label: "Active Campaigns", right: "$380" },
      { id: 4, title: "Chris Himes", label: "Average Raised", right: "$850" }]
    


    const [hasPermission, setHasPermission] = useState(null);
    const [scanned, setScanned] = useState(false);
  
    useEffect(() => {
      (async () => {
        const { status } = await BarCodeScanner.requestPermissionsAsync();
        setHasPermission(status === 'granted');
      })();
    }, []);
  

  return (
      <ScrollView style={{ backgroundColor:'#fff'}}>
    <View style={styles.container}>
    
    <View style={styles.card}>
      <Image source={{ uri: data.cover }} style={{ height:150 , width: '100%', marginBottom:25, borderRadius:10}} />
      <Text style={styles.title}>{data.title}</Text>
      <Text style={styles.raised}>{data.raised} Raised</Text>
      <View style={[{height:15, width:'100%', borderRadius: 10, borderColor: '#ddd', borderWidth:1, marginTop:10}]}>
      <Animated.View style={[{backgroundColor: "#EB7251" , borderRadius: 10, height:15, width:progressWidth }]}/>
      </View>

      <View style={styles.spacer}></View>
      <Text style={styles.sectionHeader}>Donations</Text>


      <ListItems data={donations} />

      

      <TouchableOpacity style={styles.button}><Text style={styles.buttonText}>Donate</Text></TouchableOpacity>
</View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
    </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding:0,
  },
  card: {
      padding:30
  },
  title: {
    fontSize: 30,
    fontWeight: 'bold',
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
  attendee: {
    marginTop:10,
    fontSize:25,
    marginBottom:30
  },
  camera: {
      width:width,
      height:200,
      marginBottom:50
  },
  raised: {
    marginTop:20,
    fontWeight:'bold',
    textAlign:'right'
  },
  button : {
      display:'flex',
      backgroundColor:'#f4775e',
      paddingVertical:20,
      paddingHorizontal:40,
      marginTop:15,
      alignItems:'center',
      justifyContent: 'center',
      borderRadius:10
  },
  sectionHeader: {
    fontSize: 18,
    fontWeight: 'bold',
    marginTop: 10
  },
  buttonText : {

    color : '#fff',
    fontSize:20,
    fontWeight:'bold'
  },
  spacer: {
    marginBottom: 25
  }
});
