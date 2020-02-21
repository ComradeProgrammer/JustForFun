import java.sql.*;
import java.util.HashMap;
import java.util.Scanner;

public class MainMenu {
    private Connection connection=null;
    private Statement statement=null;
    public void run() {
        connectTodatabases();
        while(true) {
            String language = selectOrCreateLanguageMenu();
            if(!language.equals("")){
                LanguageMenu l=new LanguageMenu(connection,statement,language);
                l.run();
            }
        }

    }
    protected void connectTodatabases(){
        System.out.println("欢迎使用");
        try{
            Class.forName("org.sqlite.JDBC");
            connection= DriverManager.getConnection("jdbc:sqlite:data/test.db");
            statement=connection.createStatement();
        }
        catch(Exception e){
            System.err.println( e.getClass().getName() + ": " + e.getMessage() );
            e.printStackTrace();
            System.exit(0);
        }

        System.out.println("按回车以继续");
        Scanner stdin=new Scanner(System.in);
        stdin.nextLine();
        Main.clearConsole();
    }

    protected String selectOrCreateLanguageMenu(){
        while(true) {
            int cnt = 0;
            HashMap<Integer, String> languageList = new HashMap<>();
            System.out.println("可选的语言列表");
            System.out.println("----------------------");
            try {
                ResultSet res = statement.executeQuery("select name from sqlite_master");
                while (res.next()) {
                    String language = res.getString("name");
                    if(language.charAt(0)=='_'){
                        continue;
                    }
                    languageList.put(cnt, language);
                    System.out.println(cnt + ":" + language);
                    cnt++;
                }
                if (cnt == 0) {
                    System.out.println("无可选语言");
                    System.out.println("----------------------");
                } else {
                    System.out.println("----------------------");
                    System.out.println("输入所选的语言的编号");
                }
                System.out.println("输入-1添加一门新语言,-2 删除一门语言");
                System.out.println("输入b退出");
            } catch (Exception e) {
                System.err.println(e.getClass().getName() + ": " + e.getMessage());
                e.printStackTrace();
                System.exit(0);
            }
            Scanner stdin=new Scanner(System.in);
            String input=stdin.nextLine();
            if(input.equals("b")){
                System.out.println("exit");
                System.exit(0);
            }
            int choice=0;
            try {
                choice = Integer.parseInt(input);
            }
            catch(Exception e){
                System.out.println("Invalid Input");
                return "";
            }
            Main.clearConsole();
            switch (choice){
                case -1:
                    addLanguageMenu();
                    break;
                case -2:
                    removeLanguageMenu(languageList);
                    break;
                default:
                    if(choice>=cnt){
                        System.out.println("Invalid number!");
                    }
                    else{
                        return languageList.get(choice);
                    }
            }
        }

    }
    protected void addLanguageMenu(){
        System.out.println("输入语言的英文名称");
        Scanner stdin=new Scanner(System.in);
        String newLanuage=stdin.nextLine();
        String sql="create table "+newLanuage+" ("
                +"word char(30), "
                +"partofspeech char(10),"
                +"meaning char(30),"
                +"note varchar(2000),"
                +"level int,"
                +"insertdate date"
                +");";
        String sql2="create table _"+newLanuage+"Log ("
                +"word char(30),"
                +"operation int,"
                +"updatedate date"
                +");";
        try {
            statement.executeUpdate(sql);
            statement.executeUpdate(sql2);
        }catch(Exception e){
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            e.printStackTrace();
            System.exit(0);
        }
        Main.clearConsole();

    }
    protected void removeLanguageMenu(HashMap<Integer,String>languageList){
        System.out.println("In order to protect the data, deleting language is not allowed. Sorry");
        System.out.println("Execute sql code in sqlite command line to remove a language");
        System.exit(0);
    }
}
