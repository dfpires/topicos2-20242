import React from "react";
import {
  TextStyle,
  StyleSheet,
  SafeAreaView,
  View,
  Text,
  TouchableOpacity,
  TextInput,
} from "react-native";
import { StatusBar } from "expo-status-bar";
import Constants from "expo-constants";
import { FontAwesome5 } from "@expo/vector-icons";

export default function MomoLogin() {
  return (
    <SafeAreaView style={styles.container}>
      <StatusBar style="light" />

      <View style={styles.content}>
        <View style={styles.textWrapper}>
          <Text style={styles.hiText}>Oi</Text>
          <Text style={styles.userText}>DANG QUANG HUY</Text>
          <Text style={styles.userText}>0123456789</Text>
        </View>

        <View style={styles.form}>
          {/* https://docs.expo.io/guides/icons/ */}
          <FontAwesome5 name="lock" style={styles.iconLock} />

          {/* https://reactnative.dev/docs/textinput */}
          <TextInput
            style={styles.inputPassword}
            secureTextEntry={true}
            autoFocus={true}
            placeholder="Digite a senha"
            placeholderTextColor="#929292"
          />

          {/* https://reactnative.dev/docs/touchableopacity */}
          <TouchableOpacity style={styles.buttonLogin}>
            <Text style={styles.buttonLoginText}>Conecte-se</Text>
          </TouchableOpacity>
        </View>

        <View style={styles.action}>
          <TouchableOpacity>
            <Text style={styles.userText}>ESQUECEU A SENHA</Text>
          </TouchableOpacity>

          <TouchableOpacity>
            <Text style={styles.userText}>THOÁT TÀI KHOẢN</Text>
          </TouchableOpacity>
        </View>
      </View>
    </SafeAreaView>
  );
}

const TEXT: TextStyle = {
  color: "#fff",
  textAlign: "center",
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#b0006d",
    paddingTop: Constants.statusBarHeight,
  },
  content: {
    paddingHorizontal: 30,
  },
  textWrapper: {
    marginTop: 60,
    marginBottom: 30,
  },
  hiText: {
    ...TEXT,
    fontSize: 20,
    lineHeight: 50,
    fontWeight: "bold",
  },
  userText: {
    ...TEXT,
    fontSize: 16,
    lineHeight: 30,
  },
  form: {
    marginBottom: 30,
  },
  iconLock: {
    color: "#929292",
    position: "absolute",
    fontSize: 16,
    top: 22,
    left: 22,
    zIndex: 10,
  },
  inputPassword: {
    height: 60,
    borderRadius: 30,
    paddingHorizontal: 30,
    fontSize: 20,
    color: "#929292",
    backgroundColor: "#fff",
    textAlign: "center",
    textAlignVertical: "center",
  },
  buttonLogin: {
    height: 50,
    borderRadius: 25,
    backgroundColor: "#8d015a",
    justifyContent: "center",
    marginTop: 15,
  },
  buttonLoginText: {
    ...TEXT,
  },
  action: {
    flexDirection: "row",
    justifyContent: "space-between",
  },
});
