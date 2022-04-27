import { StyleSheet , Dimensions, Touchable } from 'react-native';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { BarCodeScanner } from 'expo-barcode-scanner';
import { useEffect, useState } from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';


const { width } = Dimensions.get('window');


export default function CampaignDetails({ navigation }) {


    const [hasPermission, setHasPermission] = useState(null);
    const [scanned, setScanned] = useState(false);
  
    useEffect(() => {
      (async () => {
        const { status } = await BarCodeScanner.requestPermissionsAsync();
        setHasPermission(status === 'granted');
      })();
    }, []);
  
    const handleBarCodeScanned = ({ type, data }) => {
      setScanned(true);
      alert(`Bar code with type ${type} and data ${data} has been scanned!`);
    };
  
    if (hasPermission === null) {
      return <Text>Requesting for camera permission</Text>;
    }
    if (hasPermission === false) {
      return <Text>No access to camera</Text>;
    }

    


  return (
    <View style={styles.container}>
    
    <View style={styles.camera}>
    <BarCodeScanner
        onBarCodeScanned={scanned ? undefined : handleBarCodeScanned}
        style={StyleSheet.absoluteFillObject}
      />
    </View>
    <View style={styles.card}>
      <Text style={styles.event}>Classy Gala</Text>
      <Text style={styles.title}>General Admission</Text>
      <Text style={styles.attendee}>Omid Borjian</Text>
      <TouchableOpacity style={styles.button}><Text style={styles.buttonText}>Check In</Text></TouchableOpacity>
      <TouchableOpacity style={styles.button}><Text style={styles.buttonText}>Send Receipt</Text></TouchableOpacity>
</View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
    </View>
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
  buttonText : {

    color : '#fff',
    fontSize:20,
    fontWeight:'bold'
  }
});
