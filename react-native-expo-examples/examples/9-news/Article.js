import React from "react";
import {
  StyleSheet,
  View,
  Text,
  Image,
  TouchableOpacity,
  Linking,
  Alert,
} from "react-native";
import moment from "moment";

export default function Article({ item }) {
  const openLink = (url) => {
    // https://reactnative.dev/docs/linking
    Linking.canOpenURL(url).then((supported) => {
      if (supported) {
        Linking.openURL(url);
      } else {
        // https://reactnative.dev/docs/alert
        Alert.alert("Broken Link!");
      }
    });
  };

  return (
    <View style={styles.article}>
      {/* Caching image for better performance: https://github.com/DylanVann/react-native-fast-image */}
      <Image source={{ uri: item.url }} style={styles.articleImage} />

      <View style={{ flex: 1 }}>
        <TouchableOpacity onPress={() => openLink(item.url)}>
          <Text style={styles.articleTitle} numberOfLines={3}>
            {item.title}
          </Text>
        </TouchableOpacity>

        <Text style={styles.articlePublishedAt}>
          {moment(item.publishedAt).fromNow()}
        </Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  article: {
    flexDirection: "row",
    paddingVertical: 15,
  },
  articleImage: {
    width: 150,
    height: 85,
    resizeMode: "contain",
    marginRight: 15,
  },
  articleTitle: {
    fontSize: 18,
    fontWeight: "bold",
    marginBottom: 10,
  },
  articlePublishedAt: {
    fontSize: 14,
  },
});
