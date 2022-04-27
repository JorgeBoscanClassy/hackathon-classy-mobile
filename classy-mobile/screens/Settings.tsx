import { StyleSheet , Dimensions, Touchable } from 'react-native';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { useEffect, useState } from 'react';
import { TouchableOpacity } from 'react-native-gesture-handler';


const { width } = Dimensions.get('window');


export default function SettingsScreen() {



  return (
    <View style={styles.container}>
    
    
    <View style={styles.card}>
      <Text style={styles.event}>Hackathon</Text>
      <Text style={styles.title}>Once mobile gets funded, you'll see more cool stuff!</Text>
</View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding:0,
    justifyContent:'center'
  },
  card: {
      padding:30
  },
  title: {
    fontSize: 30,
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
  },
  event : {

  }
});
