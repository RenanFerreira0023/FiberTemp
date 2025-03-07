//+------------------------------------------------------------------+
//|                                                 RDS COPY TRADER CLI.mq5 |
//|                        Copyright 14/04/23,                     . |
//|                            Developer,         Renan D. Ferreira. |
//|                            Analyst,                  Igor Mello. |
//|                                                                  |
//|                   rdstraderprofissional@gmail.com                |
//|                                OU                                |
//|                    NOS SIGA NO INSTAGRAM [  @RDSTRADER_  ]       |
//|                                                                  |
//|                                                    RDS COPY TRADER CLI  |
//+------------------------------------------------------------------+
#property copyright "Renan Dutra Ferreira"
#property link      "https://www.mql5.com/en/users/renandutra/"
#property description "@_fdutra | appsskilldeveloper@gmail.com ";
#define VERSION "1.0"
#property version VERSION
#define NAME_BOT "RDS COPY TRADER CLI"
#include <JAson.mqh>
#include <Trade\Trade.mqh>
CTrade trade;

bool static YES   = true;
bool static NO    = false;

int static UNLOCK = 0;
int static LOCK   = 1;


enum SELECT_TYPE_TIME_ZONE
  {
   time_zone_less_12     =  -12,        // -12
   time_zone_less_11     =  -11,        // -11
   time_zone_less_10     =  -10,        // -10
   time_zone_less_9     =  -9,        // -9
   time_zone_less_8     =  -8,        // -8
   time_zone_less_7     =  -7,        // -7
   time_zone_less_6     =  -6,        // -6
   time_zone_less_5     =  -5,        // -5
   time_zone_less_4     =  -4,        // -4
   time_zone_less_3     =  -3,        // -3
   time_zone_less_2     =  -2,        // -2
   time_zone_less_1     =  -1,        // -1
   time_zone_zero       =  0,        //  0
   time_zone_max_1      = 1,           // 1
   time_zone_max_2      = 2,           // 2
   time_zone_max_3      = 3,           // 3
   time_zone_max_4      =4,            // 4
   time_zone_max_5      =5,            // 5
   time_zone_max_6      =6,            // 6
   time_zone_max_7      =7,            // 7
   time_zone_max_8      =8,            // 8
   time_zone_max_9      =9,            // 9
   time_zone_max_10      = 10,           // 10
   time_zone_max_11      =  11,          // 11
   time_zone_max_12      =  12          // 12
  };


input group                         "Login"
input string                        CHANNEL_EMAIL_LOGIN_PROTOCOL                              = "techvisionstore@gmail.com";     // Login
input string                        CHANNEL_NAME_PROTOCOL                                     = "CANAL_OURO";                // Nome do canal
input SELECT_TYPE_TIME_ZONE         TYPE_TIME_ZONE                                             = time_zone_zero;           //  Fuso horário

input double                        VOLUME_TO_SEND_ORDER                                     = 0;                       // ( 0 = Volume do servidor ) Volume a ser enviado






static int GLOBAL_LIMIT_LOOP = 10;
const string URL_SERVER = "http://192.168.1.6:8080";


int GLOBAL_AGENT_ID = -1;
int GLOBAL_RECPTOR_ID = -1;
int GLOBAL_CHANNEL_ID = -1;


string msg=""; // queue of received messages

struct Struct_Request_Copy
  {
   int               id;
   string            symbol;
   string            action_type;
   ulong             ticket;
   double            lot;
   double            target_pedding;
   double            takeprofit;
   double            stoploss;
   datetime          dt_send_order;
   int               user_agent_id;
   int               channel_id;
  };

//------------------------------------------------------------------ OnInit
int OnInit()
  {

   GLOBAL_AGENT_ID = -1;
   GLOBAL_RECPTOR_ID = -1;
   GLOBAL_CHANNEL_ID = -1;


   if(autentication_server() != INIT_SUCCEEDED)
      return INIT_FAILED;

   ArrayResize(GLOBAL_STRUCT_COPY,0);
   GLOBAL_JSON_COPY = "";

   EventSetTimer(1);
//   EventSetMillisecondTimer(100);
   return INIT_SUCCEEDED;
  }
//------------------------------------------------------------------ OnInit
void OnDeinit(const int reason)
  {
   EventKillTimer();
   Comment("");
  }

Struct_Request_Copy GLOBAL_STRUCT_COPY[];
string GLOBAL_JSON_COPY = "";
//------------------------------------------------------------------ OnInit
void OnTimer()
  {
//   Comment("\n"+NAME_BOT +" v"+VERSION+" | Port: "+Port+"  |  "+ACTION_KEY+"  | "+string(TimeCurrent())+"");


   /*


         {
            "all_copy_id":10,
            "users_receptor_id":1,
            "channel_id":1,
            "dt_send_copy":"2024-05-27 19:41:36"
         }


   */


   GLOBAL_JSON_COPY = "";
   int retCode = 0;
   string jsonRequest= "";
   string urlToken = URL_SERVER+"/Copy/Find/?";
   urlToken += "id_agent="+GLOBAL_AGENT_ID;
   urlToken += "&id_channel="+GLOBAL_CHANNEL_ID;
   urlToken += "&id_receptor="+GLOBAL_RECPTOR_ID;


   datetime timeEnd = (TimeCurrent() + (PeriodSeconds(PERIOD_M1) * 1));
   if(TYPE_TIME_ZONE != time_zone_zero)
      timeEnd += (PeriodSeconds(PERIOD_H1) * TYPE_TIME_ZONE);
   string dtEnd = TimeToString(timeEnd,TIME_DATE  | TIME_MINUTES | TIME_SECONDS);
   StringReplace(dtEnd,".","-");
   urlToken += "&end_date="+dtEnd;




   datetime timeStart = (TimeCurrent() - (PeriodSeconds(PERIOD_M1) * 10));
   if(TYPE_TIME_ZONE != time_zone_zero)
      timeStart += (PeriodSeconds(PERIOD_H1) * TYPE_TIME_ZONE);
   string dtStart = TimeToString(timeStart,TIME_DATE  | TIME_MINUTES | TIME_SECONDS);
   StringReplace(dtStart,".","-");
   urlToken += "&start_date="+dtStart;

   urlToken += "&page=0";
   urlToken += "&limit=30";

   Print("GET \n"+urlToken);
   get_webrequest(urlToken,YES,retCode,jsonRequest);
   if(retCode == 200 || retCode == 201)
     {
      GLOBAL_JSON_COPY = jsonRequest;
      print_panel_log("    Ultima Copy Recebida com sucesso! "+TimeCurrent()+"\n    Aguarde!");
     }
   else
     {
      if(retCode == 1001)
        {
         Print("Servidor fora do ar .");
         print_panel_log("Servidor fora do ar .");
        }
      else
        {
         if(retCode != 404)
           {
            CJAVal json;
            json.Deserialize(jsonRequest);
            string  msgError = json["message_error"].ToStr();
            Print("  -     "+retCode);
            Print("erro na solicitação    "+msgError);
            print_panel_log("erro na solicitação    "+msgError+"\n"+retCode);
            return;
           }
         if(retCode == 404)
           {


            print_panel_log("       Procurando Copys, aguarde        [ "+TimeCurrent()+" ]");

           }

        }
     }

   Struct_Request_Copy s_req_copy[];

   if(GLOBAL_JSON_COPY != "")
     {


      CJAVal json;
      json.Deserialize(GLOBAL_JSON_COPY);
      ArrayResize(s_req_copy,json.Size());

      for(int i = 0; i < json.Size(); i++)
        {
         s_req_copy[i].id = json[i]["id"].ToInt();
         s_req_copy[i].symbol = json[i]["symbol"].ToStr();
         s_req_copy[i].ticket = json[i]["ticket"].ToInt();

         s_req_copy[i].action_type = json[i]["action_type"].ToStr();
         s_req_copy[i].lot = json[i]["lot"].ToDbl();
         s_req_copy[i].target_pedding = json[i]["target_pedding"].ToDbl();
         s_req_copy[i].takeprofit = json[i]["takeprofit"].ToDbl();
         s_req_copy[i].stoploss = json[i]["stoploss"].ToDbl();
         s_req_copy[i].dt_send_order = StringToTime(json[i]["dt_send_order"].ToStr());

         s_req_copy[i].user_agent_id = json[i]["user_agent_id"].ToInt();
         s_req_copy[i].channel_id = json[i]["channel_id"].ToInt();


         string keyReturned = "";
         string symbolReturned = s_req_copy[i].symbol;
         string typyAction = s_req_copy[i].action_type;
         ulong tktReturned = (ulong)s_req_copy[i].ticket;


         double lotReturned = (double)s_req_copy[i].lot ;
         double priceTargetLimit = (double)s_req_copy[i].target_pedding;

         double tpReturned = (double)s_req_copy[i].takeprofit;
         double slReturned = (double)s_req_copy[i].stoploss;

         if(StringFind(symbolReturned,_Symbol,0) != -1)
           {


            if(VOLUME_TO_SEND_ORDER != 0)
               lotReturned = VOLUME_TO_SEND_ORDER;


            if(typyAction == "BUY")
               buy_market(tpReturned,slReturned,lotReturned,"#"+tktReturned);

            if(typyAction == "SELL")
               sell_market(tpReturned,slReturned,lotReturned,"#"+tktReturned);


            if(typyAction == "DELETAR")
              {
               int total=PositionsTotal();

               if(total > 0)
                  for(int cnt=0; cnt<total; cnt++)
                    {
                     string symbol = PositionGetSymbol(cnt);
                     string commet  = PositionGetString(POSITION_COMMENT);
                     ulong ticket = PositionGetInteger(POSITION_TICKET);
                     if(StringFind(commet,tktReturned,0) >= 0)
                       {
                        trade.PositionClose(ticket);
                       }
                    }
              }

            if(typyAction == "DELETAR_TUDO")
              {
               delete_all_orders_opened(tktReturned);
              }

            if(typyAction == "EDITAR")
              {
               int total=PositionsTotal();

               if(total > 0)
                  for(int cnt=0; cnt<total; cnt++)
                    {
                     string symbol = PositionGetSymbol(cnt);
                     string commet  = PositionGetString(POSITION_COMMENT);
                     ulong ticket = PositionGetInteger(POSITION_TICKET);
                     if(StringFind(commet,tktReturned,0) >= 0)
                       {
                        modify_stop_to_tkt(ticket,tpReturned,slReturned);
                       }
                    }
              }

            if(typyAction == "BUY_LIMIT" || typyAction == "BUY_STOP")
              {
               pedding_buy(priceTargetLimit,lotReturned,tpReturned,slReturned,"#"+tktReturned);
              }

            if(typyAction == "SELL_LIMIT" || typyAction == "SELL_STOP")
              {
               pedding_sell(priceTargetLimit,lotReturned,tpReturned,slReturned,"#"+tktReturned);
              }


            if(typyAction == "DEL_PEDDING")
              {
               Print("Deletar ordem  "+tktReturned);

               uint total = OrdersTotal();
               if(total > 0)
                 {
                  ulong ticket = 0;
                  for(uint i=0; i<total; i++)
                     if((ticket=OrderGetTicket(i))>0)
                       {
                        double symbol = OrderGetString(ORDER_SYMBOL);
                        double order_magic = OrderGetInteger(ORDER_MAGIC);
                        ulong ticket = OrderGetInteger(ORDER_TICKET);
                        ulong takeProfit = adjust_price(OrderGetDouble(ORDER_TP));
                        ulong stopLoss = adjust_price(OrderGetDouble(ORDER_SL));
                        string commet = OrderGetString(ORDER_COMMENT);
                        if(StringFind(commet,tktReturned,0) >= 0)
                          {
                           trade.OrderDelete(ticket);
                          }
                       }
                 }
              }
            if(typyAction == "MODIFICAR_PEDDING")
              {


               Print("MODIFICAR         ");

               uint total = OrdersTotal();
               if(total > 0)
                 {
                  ulong ticket = 0;
                  ENUM_ORDER_TYPE orderType ;
                  for(uint i=0; i<total; i++)
                     if((ticket=OrderGetTicket(i))>0)
                       {
                        double symbol = OrderGetString(ORDER_SYMBOL);
                        double order_magic = OrderGetInteger(ORDER_MAGIC);
                        ulong ticket = OrderGetInteger(ORDER_TICKET);
                        ulong takeProfit = adjust_price(OrderGetDouble(ORDER_TP));
                        ulong stopLoss = adjust_price(OrderGetDouble(ORDER_SL));
                        string commet = OrderGetString(ORDER_COMMENT);
                        if(StringFind(commet,tktReturned,0) >= 0)
                          {
                           orderType = (ENUM_ORDER_TYPE)OrderGetInteger(ORDER_TYPE);
                           trade.OrderDelete(ticket);

                           if(orderType == ORDER_TYPE_BUY || orderType == ORDER_TYPE_BUY_LIMIT ||
                              orderType == ORDER_TYPE_BUY_STOP || orderType == ORDER_TYPE_BUY_STOP_LIMIT)
                             {
                              pedding_buy(priceTargetLimit,lotReturned,tpReturned,slReturned,"#"+tktReturned);
                             }

                           if(orderType == ORDER_TYPE_SELL || orderType == ORDER_TYPE_SELL_LIMIT ||
                              orderType == ORDER_TYPE_SELL_STOP || orderType == ORDER_TYPE_SELL_STOP_LIMIT)
                             {
                              pedding_sell(priceTargetLimit,lotReturned,tpReturned,slReturned,"#"+tktReturned);
                             }

                          }
                       }
                 }
              }

            send_copy(s_req_copy[i].id);
            // criar o post
           }

        }
     }

  }






//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
string create_json_send_order(string idAllCopy)
  {

   string itemJson = "";
   itemJson += "{";
   itemJson += "\"all_copy_id\":"+idAllCopy+",";
   itemJson += "\"users_receptor_id\":"+GLOBAL_RECPTOR_ID+",";
   itemJson += "\"channel_id\":"+GLOBAL_CHANNEL_ID+",";

   string dtSendOrder = (TimeToString(TimeCurrent(),TIME_DATE  | TIME_MINUTES | TIME_SECONDS));
   StringReplace(dtSendOrder,".","-");
   itemJson += "\"dt_send_copy\":\""+dtSendOrder+"\"";

   itemJson += "}";
   return itemJson;
  }



//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void send_copy(int idAllCopy)
  {

   string jsonData = create_json_send_order(idAllCopy);


   Print("jsonData      "+jsonData);
   string msgError = "";
   int loopAuth = 0;
   do
     {
      msgError = "";
      int retCode = 0;
      string jsonToken = "";
      string urlToken = URL_SERVER+"/Copy/Reply";


      post_webrequest(urlToken,jsonData,retCode,jsonToken);

      if(retCode == 200 || retCode == 201)
        {
         msgError = "";
         Print("Confirmação da requisição realizada com sucesso !");
         loopAuth = GLOBAL_LIMIT_LOOP + 2;
         break;
        }
      else
        {
         if(retCode == 1001)
            Print("Servidor fora do ar .");
         else
           {
            CJAVal json;
            json.Deserialize(jsonToken);
            msgError = json["message_error"].ToStr();
            Print("  -     "+retCode);
            Print("erro na solicitação    "+msgError);
           }
         loopAuth++;
         Print("Tentativa de conecção :   "+loopAuth+" / "+GLOBAL_LIMIT_LOOP+"\n"+msgError);
         Sleep(700);
        }
     }
   while(loopAuth <= GLOBAL_LIMIT_LOOP);

   if(msgError != "")
     {
      Print(msgError);
      Alert(msgError);
     }



  }

//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void post_webrequest(string url, string jsonData, int &retCode, string &requestJson)
  {
   int jsonDataSize = StringLen(jsonData);
   uchar jsonDataChar[];
   StringToCharArray(jsonData, jsonDataChar, 0,jsonDataSize,CP_UTF8);

   string headers = "Authorization:"+GLOBAL_KEY_HEADER;

   uchar serverResult[];
   string serverHeaders;
   retCode = WebRequest("POST", url,headers,500,jsonDataChar,  serverResult, serverHeaders);
   if(retCode == 200 || retCode == 201)
      requestJson = CharArrayToString(serverResult,0,ArraySize(serverResult), CP_UTF8);
   else
      if(retCode == 1001)
        {
         requestJson = ("Servidor não encontrado");
        }
      else
         requestJson = CharArrayToString(serverResult,0,ArraySize(serverResult), CP_UTF8);
  }

//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void get_webrequest(string url, bool headersUsing,int &retCode, string &requestJson)
  {
   string jsonData = "";
   int jsonDataSize = 0;
   uchar jsonDataChar[];
   StringToCharArray(jsonData, jsonDataChar, 0,jsonDataSize,CP_UTF8);

   string headers = "";
   if(headersUsing == YES)
      headers = "Authorization:"+GLOBAL_KEY_HEADER;
   else
      headers = "";

   uchar serverResult[];
   string serverHeaders;
   retCode = WebRequest("GET", url,headers,500,jsonDataChar,  serverResult, serverHeaders);
   if(retCode == 200 || retCode == 201)
      requestJson = CharArrayToString(serverResult,0,ArraySize(serverResult), CP_UTF8);
   else
      if(retCode == 1001)
        {
         requestJson = ("Servidor não encontrado");
        }
      else
         requestJson = CharArrayToString(serverResult,0,ArraySize(serverResult), CP_UTF8);

  }


//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void print_panel_log(string msg)
  {
   string textComment = "";
   textComment  += "\n########################################";
   textComment  += "\n############## META COPY 5 ################";
   textComment  += "\n########################################";
   textComment  += "\n\n     "+msg;
   textComment  += "\n\n########################################";
   textComment  += "\n######## NÃO FECHAR ESTE PROGRAMA ##########";
   textComment  += "\n########################################";
   Comment(textComment);
  }

//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
int autentication_server()
  {
   print_panel_log("Autenticando com o servidor. \n     Aguarde !");
   string msgError = "";
   int loopAuth = 0;
   do
     {
      msgError = "";
      int retCode = 0;
      string jsonToken = "";
      string urlToken = URL_SERVER+"/Repector/Auth/?login="+CHANNEL_EMAIL_LOGIN_PROTOCOL;

      get_webrequest(urlToken,NO,retCode,jsonToken);

      if(retCode == 200 || retCode == 201)
        {
         msgError = "";
         //         Print("Token gerado com sucesso !");
         CJAVal json;
         json.Deserialize(jsonToken);
         GLOBAL_KEY_HEADER = json["token"].ToStr();
         loopAuth = GLOBAL_LIMIT_LOOP + 2;
         //         Print("id receptor :  "+json["id"].ToInt());

         break;
        }
      else
        {
         if(retCode == 1001)
            msgError += ("Servidor fora do ar .");
         else
           {
            CJAVal json;
            json.Deserialize(jsonToken);
            msgError = json["message_error"].ToStr();
            Print("  -     "+retCode);
            Print("erro na solicitação    "+msgError);
           }
         loopAuth++;
         Print("Tentativa de conecção :   "+loopAuth+" / "+GLOBAL_LIMIT_LOOP);
         print_panel_log("Tentativa de conecção :   "+loopAuth+" / "+GLOBAL_LIMIT_LOOP +"\n"+msgError);
         Sleep(700);
        }
     }
   while(loopAuth <= GLOBAL_LIMIT_LOOP);


   if(msgError != "")
     {
      Print(msgError);
      Alert(msgError);
     }

   msgError = "";


   if(GLOBAL_KEY_HEADER != "")
     {
      //      Print(GLOBAL_KEY_HEADER);
      msgError = "";
      loopAuth = 0;

      do
        {
         msgError = "";
         int retCode = 0;
         string jsonRequest = "";
         string urlToken = URL_SERVER+"/Receptor/Login/mt5/?login="+CHANNEL_EMAIL_LOGIN_PROTOCOL+"&channel="+CHANNEL_NAME_PROTOCOL;
         get_webrequest(urlToken,YES,retCode,jsonRequest);

         if(retCode == 200 || retCode == 201)
           {
            Print("Login Realizado com sucesso !. Aguarde!");
            print_panel_log("Login Realizado com sucesso !. Aguarde!");

            msgError = "";
            CJAVal json;
            json.Deserialize(jsonRequest);
            GLOBAL_AGENT_ID = json["user_agent_id"].ToInt();
            GLOBAL_RECPTOR_ID = json["user_receptor_id"].ToInt();
            GLOBAL_CHANNEL_ID = json["channel_id"].ToInt();
            //            Print("GLOBAL_AGENT_ID  "+GLOBAL_AGENT_ID);
            //            Print("GLOBAL_RECPTOR_ID  "+GLOBAL_RECPTOR_ID);
            //            Print("GLOBAL_CHANNEL_ID  "+GLOBAL_CHANNEL_ID);



            loopAuth = GLOBAL_LIMIT_LOOP + 2;


            break;
           }
         else
           {
            if(retCode == 1001)
               msgError += ("Servidor fora do ar .");
            else
              {
               CJAVal json;
               json.Deserialize(jsonRequest);
               msgError = json["message_error"].ToStr();
               Print("  -     "+retCode);
               Print("erro na solicitação    "+msgError);
              }
            loopAuth++;
            Print("Tentativa de conecção :   "+loopAuth+" / "+GLOBAL_LIMIT_LOOP);
            print_panel_log("Tentativa de conecção :   "+loopAuth+" / "+GLOBAL_LIMIT_LOOP+"\n"+msgError);
            Sleep(700);
           }
        }
      while(loopAuth <= GLOBAL_LIMIT_LOOP);
     }

   if(msgError != "")
     {
      Print(msgError);
      Alert(msgError);
      print_panel_log(msgError);
     }

   if(GLOBAL_KEY_HEADER == "" || GLOBAL_AGENT_ID == -1  || GLOBAL_RECPTOR_ID == -1 || GLOBAL_CHANNEL_ID == -1)
      return INIT_FAILED;
   else
     {
      KEY_FLOW_GLOBAL = UNLOCK;

      return INIT_SUCCEEDED;
     }
  }





string GLOBAL_KEY_HEADER = "";
int KEY_FLOW_GLOBAL = LOCK;


//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void pedding_buy(double targetBuy,double lot, double targetBuyTakeProfit,double targetBuyStopLoss,string comment)
  {
   double ask = SymbolInfoDouble(_Symbol,   SYMBOL_ASK);

   if(targetBuy < ask)
      trade.BuyLimit(lot,
                     adjust_price(targetBuy),
                     _Symbol,
                     adjust_price(targetBuyStopLoss),
                     adjust_price(targetBuyTakeProfit),
                     ORDER_TIME_GTC,0,comment);
   if(targetBuy > ask)
      trade.BuyStop(lot,
                    adjust_price(targetBuy),
                    _Symbol,
                    adjust_price(targetBuyStopLoss),
                    adjust_price(targetBuyTakeProfit),
                    ORDER_TIME_GTC,0,comment);


  }

//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void pedding_sell(double targetSell,double lot, double targetSellTakeProfit,double targetSellStopLoss,string comment)
  {
   double bid = SymbolInfoDouble(_Symbol,   SYMBOL_BID);
   if(targetSell > bid)
      trade.SellLimit(lot,
                      adjust_price(targetSell),
                      _Symbol,
                      adjust_price(targetSellStopLoss),
                      adjust_price(targetSellTakeProfit),
                      ORDER_TIME_GTC,0,comment);
   if(targetSell < bid)
      trade.SellStop(lot,
                     adjust_price(targetSell),
                     _Symbol,
                     adjust_price(targetSellStopLoss),
                     adjust_price(targetSellTakeProfit),
                     ORDER_TIME_GTC,0,comment);

  }


//+------------------------------------------------------------------+
//|                                                                  |
//| fix the number according to the decimal places of the asset      |
//|                                                                  |
//+------------------------------------------------------------------+
double adjust_price(double price)
  {
   double tickSize = SymbolInfoDouble(_Symbol, SYMBOL_TRADE_TICK_SIZE);
   return(MathRound((price)/ tickSize) * tickSize);
  }

//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
bool modify_stop_to_tkt(ulong order, double takeprofit,double stop)
  {

   bool lockedLoop = true;
   int countTry = 1;

   while(lockedLoop == true)
     {
      bool ok = trade.PositionModify(order, adjust_price(stop),adjust_price(takeprofit)) ;
      if(!ok)
        {
         int errorCode = GetLastError();
         ResetLastError();
        }
      ulong retCode = trade.ResultRetcode();
      if(retCode == 10009 || retCode == 10008)
        {
         lockedLoop = false;
        }

      Print(" order  "+order+"  modify_stop_to_tkt | RESULT RET CODE :  "+retCode + "   takeprofit "+takeprofit+"     stop "+stop);

      if(countTry > 15)
        {
         lockedLoop = false;
         return false;
        }
      else
        {
         Sleep(1500);
         countTry++;
        }
     }

   return true;
  }





//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void delete_all_orders_opened(string tktReturn2ed)
  {

   int total = PositionsTotal();

   if(total == 0)
      return;

   double arrayTktForDell[];
   ArrayFree(arrayTktForDell);
   ArrayResize(arrayTktForDell,true);
   ArrayResize(arrayTktForDell,0);
   ArrayPrint(arrayTktForDell);
   for(int i = 0; i < total; i++)
     {
      ulong ticket = PositionGetTicket(i);
      string symbol = PositionGetSymbol(i);
      ulong magic = PositionGetInteger(POSITION_MAGIC);
      string commet = PositionGetString(POSITION_COMMENT);
      //      if(StringFind(commet,tktReturned,0) >= 0)
      if(symbol == _Symbol)
        {
         ArrayResize(arrayTktForDell,ArraySize(arrayTktForDell)+1);
         arrayTktForDell[ArraySize(arrayTktForDell)-1] = ticket;
        }
     }

   if(ArraySize(arrayTktForDell) == 0)
      return; // tem ordens abertas mais nenhuma deste bot

   int numOperationOpened = ArraySize(arrayTktForDell);

   int loopCount = 0;
   while(loopCount < ArraySize(arrayTktForDell))
     {
      ulong tkt = arrayTktForDell[loopCount];

      CPositionInfo myPosition;
      if(myPosition.SelectByTicket(tkt) == true)
         if(trade.PositionClose(tkt) == true)
            loopCount++;
     };
  }



//+------------------------------------------------------------------+
//|                                                                  |
//+------------------------------------------------------------------+
void find_text(string text, string phrase, int &posStart, int &posY)
  {
// encontra a posição da frase no texto
   int startPos = StringFind(text, phrase, 0);
   int endPos = startPos + StringLen(phrase) - 1;

// exibe a posição numérica e a posição de início e fim da frase
   if(startPos >= 0)
     {
      Print("Frase encontrada na posição numérica ", startPos, " (", startPos + 1, "-", endPos + 1, ")");
      posStart = startPos + 1;
      posY = endPos + 1;
     }
   else
     {
      Print("A frase não foi encontrada no texto.");
     }
  }



//+------------------------------------------------------------------+
//|                                                                  |
//|  send BUY to market                                              |
//|                                                                  |
//+------------------------------------------------------------------+
ulong buy_market(double takeprofit,double stoploss, double lots, string comment)
  {
   double ask = SymbolInfoDouble(_Symbol,   SYMBOL_ASK);
   if(lots < 0)
      lots = +lots;

   bool ok = trade.Buy(lots, _Symbol,ask, stoploss, takeprofit,comment);
   if(!ok)
     {
      int errorCode = GetLastError();
      Print("lots    "+lots+"   BuyMarket : "+errorCode+"         |        ResultRetcode :  "+trade.ResultRetcode());
      ResetLastError();
      return -1;
     }

   Print("\n===== A MERDADO COMPRA | RESULT RET CODE :  "+trade.ResultRetcode());
   Print("LOTE ENVIADO  :  "+lots);
   ulong order = trade.ResultOrder();

   Print("TKT OFERTA : "+order);

   return order;
  }

//+------------------------------------------------------------------+
//|                                                                  |
//|  send SELL to market                                             |
//|                                                                  |
//+------------------------------------------------------------------+
ulong sell_market(double takeprofit,double stoploss, double lots, string comment)
  {
   double bid = SymbolInfoDouble(_Symbol,   SYMBOL_BID);
   if(lots < 0)
      lots = +lots;

   bool ok = trade.Sell(lots, _Symbol,bid, stoploss, takeprofit,comment);
   if(!ok)
     {
      int errorCode = GetLastError();
      Print("lots    "+lots+"    SellMarket : "+errorCode+"         |        ResultRetcode :  "+trade.ResultRetcode());
      ResetLastError();
      return -1;
     }

   Print("\n===== A MERDADO VENDA | RESULT RET CODE :  "+trade.ResultRetcode());
   Print("LOTE ENVIADO  :  "+lots);
   ulong order = trade.ResultOrder();

   Print("TKT OFERTA : "+order);

   return order;
  }

//+------------------------------------------------------------------+

//+------------------------------------------------------------------+
