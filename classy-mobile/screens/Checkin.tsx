import { StyleSheet, SafeAreaView, FlatList } from 'react-native';
import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { ListItem } from '../components/ListItems'
import { Ionicons } from '@expo/vector-icons'; 
import { TouchableOpacity } from 'react-native-gesture-handler';


export default function CheckinScreen({ navigation }) {

    const DATA = [
        {
          id: 1,
          title: 'Omid Borjian',
          label : 'Registered 2 hours ago'
        },
        {
          id: 2,
          title: 'Emad Borjian',
          label : 'Registered 2 hours ago'
        },
        {
          id: 3,
          title: 'Chris Himes',
          label : 'Registered 2 hours ago'
        },
        {
            id: 3,
            title: 'Shantanu Bose',
            label : 'Registered 2 hours ago'
        },
        {
            id: 3,
            title: 'Tammen Bruccoleri',
            label : 'Registered 2 hours ago'
        },
      ];

      const renderItem = ({ item , navigation }) => (
        <View key={item.id} style={{paddingLeft:15}}>
            <ListItem data={item} onPress={() => { navigation.navigate('Details') }} />
        </View>
      );

      
    
  return (
    <View style={styles.container}>
        <View style={{ flexDirection:'row', alignItems:'center' }}>
      <Text style={styles.title}>Attendees</Text>
      <TouchableOpacity 
      onPress={() => {
          navigation.navigate("Scanner")
      }}
      style={{ marginRight:23 }}><Ionicons name="qr-code" size={40} color="#f4775e" /></TouchableOpacity>
      </View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <SafeAreaView style={styles.container}>
      <FlatList
        data={DATA}
        renderItem={renderItem}
        keyExtractor={item => item.id}
      />
    </SafeAreaView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  title: {
    marginRight:'auto',
    fontSize: 40,
    paddingLeft:15,
    paddingTop:20,
    paddingBottom:15,
    fontWeight: 'bold',
  },
});
