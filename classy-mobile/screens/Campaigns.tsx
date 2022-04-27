import { StyleSheet, SafeAreaView, FlatList } from 'react-native';
import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { ListItem } from '../components/ListItems'
import { TouchableOpacity } from 'react-native-gesture-handler';
import {useState} from 'react'


export default function AttendessScreen({ navigation }) {

    const [ isFetching, setIsFetching ] = useState(false);

    const DATA = [
        {
          id: 1,
          title: 'Fight for Clean Air',
          label : '567 Supporters | $543,00 Raised'
        },
        {
          id: 2,
          title: 'Give Now to Help Families',
          label : '145 Supporters | $143,00 Raised'
        },
        {
          id: 3,
          title: 'Sacred Heart’s Fun Run & Pantry Drive',
          label : '564 Supporters | $643,00 Raised'
        },
        {
          id: 4,
          title: '2022 End of Year Giving',
          label : '564 Supporters | $943,00 Raised'
        },
        {
          id: 5,
          title: 'Childhood Cancer Awareness Month',
          label : '564 Supporters | $52,00 Raised'
        }
      ];

      const renderItem = ({ item  }) => (
        <View key={item.id} style={{paddingLeft:15}}>
            <ListItem data={item} onPress={() => { navigation.navigate('Campaign') }} />
        </View>
      );


      const onRefresh = () => {

      }

      
    
  return (
    <View style={styles.container}>
      <View onPress={() => { navigation.navigate('Campaign') }}><Text style={styles.title}>Campaigns</Text></View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <SafeAreaView style={styles.container}>
      <FlatList
        data={DATA}
        renderItem={renderItem}
        onRefresh={onRefresh}
        refreshing={isFetching}
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
    fontSize: 30,
    paddingLeft:15,
    paddingTop:20,
    paddingBottom:15,
    fontWeight: 'bold',
  },
});
